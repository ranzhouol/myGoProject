package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()
var RedisServer = InitRedis()

func Init() *gorm.DB {
	// 连接数据库
	dsn := "root:ranzhou@tcp(192.168.239.100:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init error: ", err)
	}

	// 建表
	if err := db.AutoMigrate(
		&SubmmitBasic{},
		&CategoryBasic{},
		&ProblemBasic{},
		&UserBasic{},
	); err != nil {
		log.Println("gorm AutoMigrate error: ", err)
	}
	return db
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.239.100:6379",
		Password: "ranzhou",
		DB:       0, // use default DB
	})
}
