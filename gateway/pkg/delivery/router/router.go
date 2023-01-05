package router

import (
	"github.com/gin-gonic/gin"

	"github.com/marceloaguero/go-auth-nats-gateway/gateway/pkg/delivery/users"
)

type router struct {
	usersDelivery users.Delivery
}

func NewRouter(usersDelivery users.Delivery, pathPrefix string) (*router, error) {
	router := &router{
		usersDelivery: usersDelivery,
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users := r.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "usuarios",
			})
			users.POST("/", router.usersDelivery.Create)
		})
	}

	err := r.Run()
	if err != nil {
		return nil, err
	}

	return router, nil
}
