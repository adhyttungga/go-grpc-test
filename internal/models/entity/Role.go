package entity

type Role struct {
	Id          int `gorm:"type:int;primaryKey" json:"id"`
	RoleRightId int `gorm:"type:int" json:"role_right_id"`
}
