package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/wahyudibo/go-casbin-example/internal/authz"
	m "github.com/wahyudibo/go-casbin-example/internal/middlewares"
)

type Router struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Router {
	return &Router{
		db: db,
	}
}

func (r *Router) Run() {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// load authorization enforcer
	authzenf, err := authz.New(r.db)
	if err != nil {
		log.Fatalf("failed to initialize authz: %+v\n", err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/authorized-only",
		m.Authz(authzenf, m.ViewOpportunityMarketplacePolicy),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "only authorized user is allowed here",
			})
		})

	router.Run()
}
