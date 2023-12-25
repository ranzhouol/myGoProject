package models

import "gorm.io/gorm"

/*
问题表和分类表的关联表
*/
type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `gorm:"column:problem_id;type:varchar(36);" json:"problem_id"`  //问题的id
	CategoryId    string         `gorm:"column:category_id;type:varchar(36)" json:"category_id"` //分类的id
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`                   //关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
