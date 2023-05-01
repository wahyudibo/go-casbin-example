package mysql

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migrationsql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/wahyudibo/go-casbin-example/internal/config"
)

type Client struct {
	config *config.Config
	conn   *gorm.DB
}

func New(cfg *config.Config) (*Client, error) {
	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{config: cfg, conn: conn}, nil
}

func connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (c *Client) Migrate() error {
	db, err := c.conn.DB()
	if err != nil {
		return err
	}

	driver, err := migrationsql.WithInstance(db, &migrationsql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		c.config.DatabaseMigrationFilePath,
		c.config.Database.Name,
		driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info("no schema changes to apply")
			return nil
		}
	}

	return err
}

func (c *Client) GetConnection() *gorm.DB {
	return c.conn
}
