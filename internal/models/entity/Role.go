package entity

type Role struct {
	Id   int64  `gorm:"type:bigint;primaryKey" json:"id"`
	Name string `gorm:"type:varchar" json:"name"`
}
