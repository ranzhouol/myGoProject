package models

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var RedisServer = InitRedis()

func Init() {
	// 连接数据库
	dsn := "root:ranzhou@tcp(192.168.239.100:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init error: ", err)
		return
	}

	if DB == nil {
		log.Println("DB is nil")
		return
	}

	fmt.Println("连接数据库成功", DB.Name())
	// 建表
	registerTables()
}

func registerTables() {
	if err := DB.AutoMigrate(
		&SubmmitBasic{},
		&CategoryBasic{},
		&UserBasic{},
	); err != nil {
		log.Fatalf("Create Db tables Error: %v", err.Error())
	}
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.239.100:6379",
		Password: "ranzhou",
		DB:       0, // use default DB
	})
}
