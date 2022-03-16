package user

import (
	"gintest/common/params"
	"gintest/config"
	"gintest/models"
	"gintest/utils"
	"strconv"

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

func ListHandler(c *gin.Context) *utils.APIException {
	user := &models.User{
		Name: "a",
	}
	users, err := user.List()
	if err != nil {
		return utils.NewAPIException(config.ERROR_LIST_USER_FAIL, err)
	}
	return utils.SuccResp("", struct {
		Users *[]models.User `json:"users"`
	}{Users: users})
}

func OneHandler(c *gin.Context) *utils.APIException {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ParameterError("id is invalid", err)
	}
	user := &models.User{
		ID: id,
	}
	err = user.One()
	if err != nil {
		return utils.NewAPIException(config.ERROR_ONE_USER_FAIL, err)
	}
	return utils.SuccResp("", struct {
		User *models.User `json:"user"`
	}{User: user})
}

// TODO
func ModifyHandler(c *gin.Context) *utils.APIException {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ParameterError("id is invalid", err)
	}
	user := &models.User{
		ID: id,
	}
	err = user.One()
	if err != nil {
		return utils.NewAPIException(config.ERROR_ONE_USER_FAIL, err)
	}

	var params params.ParamsModifyUser
	err = c.ShouldBind(&params)
	if err != nil {
		return utils.ParameterError("parameter parsing failed", err)
	}

	err = user.Update(params)
	if err != nil {
		return utils.NewAPIException(config.ERROR_MOD_USER_FAIL, nil)
	}

	return utils.SuccResp("", struct {
		Res bool `json:"res"`
	}{Res: true})
}

func DeleteHandler(c *gin.Context) *utils.APIException {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ParameterError("id is invalid", err)
	}
	user := &models.User{
		ID: id,
	}
	rows := user.Delete()
	if rows < 1 {
		return utils.NewAPIException(config.ERROR_DEL_USER_FAIL, nil)
	}
	return utils.SuccResp("", struct {
		Rows int64 `json:"rows"`
	}{Rows: rows})
}
