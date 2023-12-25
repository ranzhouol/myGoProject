package models

import (
	"gorm.io/gorm"
)

type SubmmitBasic struct {
	gorm.Model      `gorm:"table:submmit_basic"`
	Identity        string        `gorm:"column:identity;type:varchar(36);" json:"identity"`                 //提交表的唯一标识
	ProblemIdentity string        `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` //问题表的唯一标识
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity"`                   //关联问题表
	UserIdentity    string        `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       //用户表的唯一标识
	UserBasic       *UserBasic    `gorm:"foreignKey:identity;references:user_identity"`                      //关联用户表
	Path            string        `gorm:"column:path;type:varchar(255);" json:"path"`                        //代码存放路径
	Status          int           `gorm:"column:status;type:tinyint(1);" json:"status"`                      //-1-待判断,1-答案正确,2-答案错误,3-运行超时,4-运行超内存
}

func (table *SubmmitBasic) TableName() string {
	return "submmit_basic"
}

// GetSubmitList 获取提交列表
func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm.DB {
	// Preload 预热，相当于外键表连接
	tx := DB.Model(new(SubmmitBasic)).Preload("ProblemBasic").Preload("UserBasic")

	// 条件查询
	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}

	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}

	if status != 0 {
		tx.Where("status = ?", status)
	}
	return tx
}
