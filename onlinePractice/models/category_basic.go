package models

type CategoryBasic struct {
	//gorm.Model
	model
	Identity     string         `gorm:"column:identity;type:varchar(36);" json:"identity"` //分类的唯一标识
	Name         string         `gorm:"column:name;type:varchar(100);" json:"name"`        //分类名
	ParentId     int            `gorm:"column:parent_id;type:int;" json:"parent_id"`       //父级id
	ProblemBasic []ProblemBasic `gorm:"many2many:problem_category;"`
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}
