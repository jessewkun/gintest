package user

import (
	"gintest/common/params"
	"gintest/config"
	"gintest/models"
	"gintest/utils"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) *utils.APIException {
	var (
		params params.ParamsAddUser
		err    error
	)

	err = c.ShouldBind(&params)
	if err != nil {
		return utils.ParameterError("parameter parsing failed", err)
	}

	user := models.User{}
	id, err := user.Add(&params)
	if err != nil {
		return utils.NewAPIException(config.ERROR_ADD_USER_FAIL, nil)
	}

	return utils.SuccResp("", struct {
		Id int `json:"id"`
	}{Id: id})
}
