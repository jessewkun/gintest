package params

type ParamsAddUser struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
	Sex  int    `form:"sex" json:"sex" binding:"required,oneof=1 2"`
}
