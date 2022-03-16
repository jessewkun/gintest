package routers

import (
	"gintest/app/user"
	"gintest/utils"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) *gin.Engine {
	r.Use(TraceId(), LogTrace())
	r.NoMethod(utils.Wrapper(HandleNotFound))
	r.NoRoute(utils.Wrapper(HandleNotFound))

	v1User := r.Group("v1")
	v1User.GET("/users", utils.Wrapper(user.ListHandler))
	v1User.POST("/users", utils.Wrapper(user.AddHandler))
	v1User.GET("/users/:id", utils.Wrapper(user.OneHandler))
	v1User.PUT("/users/:id", utils.Wrapper(user.ModifyHandler))
	v1User.DELETE("/users/:id", utils.Wrapper(user.DeleteHandler))

	return r
}

func HandleNotFound(c *gin.Context) *utils.APIException {
	return utils.NotFound()
}
