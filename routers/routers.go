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

	v1User := r.Group("v1/user")
	v1User.POST("/add", utils.Wrapper(user.AddHandler))

	return r
}

func HandleNotFound(c *gin.Context) *utils.APIException {
	return utils.NotFound()
}
