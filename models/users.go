package models

import (
	"gintest/common/params"
	"gintest/utils"

	"github.com/jinzhu/copier"
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

func (u *User) Delete() (int64, error) {
	tx := utils.DB.Where(u).First(u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (u *User) Update(params params.ParamsModifyUser) error {
	// update 的时候 model 和 Updates 必须是相同的结构体
	var updates User
	copier.Copy(&updates, &params)
	if err := utils.DB.Model(u).Updates(&updates).Error; err != nil {
		return err
	}
	return nil
}
