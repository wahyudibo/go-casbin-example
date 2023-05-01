package authz

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"github.com/wahyudibo/go-casbin-example/internal/database/models"
)

const (
	rbacModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m =  r.obj == p.obj && g(r.sub, p.sub) && r.act == p.act
`

	rbacTableName = "permissions__users"
)

func New(db *gorm.DB) (*casbin.CachedEnforcer, error) {
	gormadapter.TurnOffAutoMigrate(db)

	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, new(models.PermissionUser), rbacTableName)
	if err != nil {
		return nil, err
	}

	model, err := casbinmodel.NewModelFromString(rbacModel)
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewCachedEnforcer(model, adapter)
	if err != nil {
		return nil, err
	}

	enforcer.LoadPolicy()

	return enforcer, nil
}
