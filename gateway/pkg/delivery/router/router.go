package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

type router struct {
	eng *gin.Engine
	nc  *nats.Conn
}

func newRouter(eng *gin.Engine, nc *nats.Conn) *router {
	return &router{
		eng: eng,
		nc:  nc,
	}
}

func NewRouter(pathPrefix, natsURLs string) (*router, error) {
	nc, err := nats.Connect(natsURLs)
	if err != nil {
		return nil, err
	}
	defer nc.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	usrs := r.Group("/users")
	{
		usrs.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "usuarios",
			})
			//usrs.POST("/", users.Create)
		})
	}

	err = r.Run()
	if err != nil {
		return nil, err
	}

	router := newRouter(r, nc)
	return router, nil
}

func (r *router) Drain() {
	r.nc.Drain()
}
