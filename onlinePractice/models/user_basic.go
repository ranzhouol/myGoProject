package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"column:identity;type:varchar(36);" json:"identity"` //用户的唯一标识
	Name             string `gorm:"column:name;type:varchar(100);" json:"name"`        //用户名
	Password         string `gorm:"column:password;type:varchar(32);" json:"password"` //密码
	Phone            string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Email            string `gorm:"column:email;type:varchar(100);" json:"email"`
	FinishProblemNum int64  `gorm:"column:finish_problem_num;type:int(11);" json:"finish_problem_num"` // 完成问题个数
	SubmitNum        int64  `gorm:"column:submit_num;type:int(11);" json:"submit_num"`                 //提交次数
}

// 定义表名，gorm默认创建的表名会加s
func (table *UserBasic) TableName() string {
	return "user_basic"
}
