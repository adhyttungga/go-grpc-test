package entity

type User struct {
	Id         int64  `gorm:"type:bigint;primaryKey" json:"id"`
	RoleId     int64  `gorm:"type:bigint" json:"role_id"`
	Name       string `gorm:"type:varchar" json:"name"`
	Password   string `gorm:"type:varchar" json:"password"`
	Email      string `gorm:"type:varchar;unique;not null" json:"email"`
	LastAccess int64  `gorm:"type:bigint" json:"last_access"`
}
