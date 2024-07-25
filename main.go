package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义数据库模型
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	dsn := "host=localhost user=root password=root dbname=demo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to get DB instance: ", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Failed to close database connection: ", err)
		}
	}()

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 使用 GORM 进行数据库操作
	// 插入数据
	db.Create(&User{
		Name:      "Alice1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	db.Create(&User{
		Name:      "Alice1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	// 查询数据
	var users []User
	db.Find(&users)
	log.Println("Users:", users)
	time.Sleep(time.Second * 10)
	// 关闭数据库连接

}
