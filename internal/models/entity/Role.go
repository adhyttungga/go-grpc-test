package entity

type Role struct {
	Id          int    `gorm:"type:int;primaryKey" json:"id"`
	Name        string `gorm:"type:varchar" json:"name"`
	RoleRightId int    `gorm:"type:int" json:"role_right_id"`
}
