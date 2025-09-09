package entity

type RoleRight struct {
	Id      int    `gorm:"type:int;primaryKey" json:"id"`
	Section string `gorm:"type:varchar" json:"section"`
	Route   string `gorm:"type:varchar" json:"route"`
	RCreate int    `gorm:"type:int" json:"r_create"`
	RRead   int    `gorm:"type:int" json:"r_read"`
	RUpdate int    `gorm:"type:int" json:"r_Update"`
	RDelete int    `gorm:"type:int" json:"r_delete"`
}
