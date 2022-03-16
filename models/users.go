package models

import (
	"gintest/common/params"
	"gintest/utils"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

// 获取表名
func (u *User) TableName() string {
	return "user"
}

// 添加用户
func (u *User) Add(params *params.ParamsAddUser) (int, error) {
	data := User{
		Name: params.Name,
		Age:  params.Age,
		Sex:  params.Sex,
	}
	result := utils.GetDB().Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.ID, nil
}
