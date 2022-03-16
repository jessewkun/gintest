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

func (u *User) Add(params *params.ParamsAddUser) (int, error) {
	data := User{
		Name: params.Name,
		Age:  params.Age,
		Sex:  params.Sex,
	}
	result := utils.DB.Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.ID, nil
}

func (u *User) List() (*[]User, error) {
	users := &[]User{}
	err := utils.DB.Where(u).Find(users).Error
	if err != nil {
		return &[]User{}, err
	}
	return users, nil
}

func (u *User) One() error {
	err := utils.DB.Where(u).First(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() int64 {
	return utils.DB.Where(u).Delete(u).RowsAffected
}

func (u *User) Update(params params.ParamsModifyUser) error {
	if err := utils.DB.Model(u).Updates(params).Error; err != nil {
		return err
	}
	return nil
}
