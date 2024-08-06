package model

type Project struct {
	Id          int64  `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	UserId      int64  `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
}

func (p *Project) TableName() string {
	return "ezp_bd_easybi_project"
}
