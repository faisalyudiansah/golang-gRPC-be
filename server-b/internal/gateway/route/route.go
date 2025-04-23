package route

import (
	controllers "server/internal/controller"

	"github.com/gin-gonic/gin"
)

func ExampleControllerRoute(c *controllers.ExampleController, r *gin.Engine) {
	g := r.Group("")
	{
		g.GET("/example", c.Get)
	}
}

func UserControllerRoute(c *controllers.UserController, r *gin.Engine) {
	g := r.Group("/users")
	{
		g.GET("/:id", c.Get)
	}
}
