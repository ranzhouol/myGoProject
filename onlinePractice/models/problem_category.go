package models

/*
问题表和分类表的关联表
*/
type ProblemCategory struct {
	//gorm.Model
	model
	ProblemId     uint          `gorm:"column:problem_id;type:int;" json:"problem_id"`   //问题的id
	CategoryId    uint          `gorm:"column:category_id;type:int;" json:"category_id"` //分类的id
	CategoryBasic CategoryBasic `gorm:"foreignKey:category_id"`                          //关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
