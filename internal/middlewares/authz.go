package middlewares

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Authz is middleware to enforce authorization.
func Authz(e *casbin.CachedEnforcer, policy Policy) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := "123"

		permissions, err := e.GetImplicitPermissionsForUser(userID)
		if err != nil {
			log.Errorf("failed when loading permissions for user %s: %+v\n", userID, err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		var isUserAuthorized bool
		for _, p := range permissions {
			if p[1] == policy.Resource && p[2] == policy.Action {
				isUserAuthorized = true
			}
		}

		if !isUserAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}
