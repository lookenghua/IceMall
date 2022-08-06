package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "ice-mall/schema"
	"log"
)

var client *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	dsn := "root:123456@tcp(localhost:3306)/ice_mall?parseTime=true"
	var err error
	client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	err1 := client.AutoMigrate(&User{})
	if err1 != nil {
		log.Fatalf("迁移数据库失败")
		return
	}
	SeedDatabase()
}

// SeedDatabase 初始化数据库数据
func SeedDatabase() {
	fmt.Printf("===开始初始化数据===\n")
	// 设置验证码发送间隔
	{
		/*n, err := client.Model.Query().Where(systemconfig.RuleEQ("SEND_CODE_TIME")).Count(context.Background())
		if err != nil && n == 0 {
			_, e := client.SystemConfig.Create().SetRule("SEND_CODE_TIME").SetValue("5").SetRemark("发送验证码间隔").Save(context.Background())
			if e != nil {
				fmt.Printf("%v\n", e)
				return
			}
		}*/
	}

	fmt.Println("===初始化数据成功===")
}
