package model

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "ice-mall/schema"
	"ice-mall/schema/user"
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
	err1 := client.AutoMigrate(&User{}, &Product{}, &Annex{})
	if err1 != nil {
		log.Fatalf("迁移数据库失败")
		return
	}
	SeedDatabase()
}

// SeedDatabase 初始化数据库数据
func SeedDatabase() {
	fmt.Printf("===开始初始化数据===\n")
	// 初始化admin账号
	{
		u := User{}
		result := client.First(&u, "username = ?", "admin")
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("查询admin账号失败")
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("初始化admin账号")
			// 创建admin
			result := client.Create(&User{
				Username: "admin",
				Password: "a8c9b651cbd5f970ec397d95f8054ec4",
				Role:     user.RoleADMIN,
			})
			if result.Error != nil {
				fmt.Println("创建admin失败")
			} else {
				fmt.Println("创建admin成功")
			}
		} else {
			fmt.Println("admin账号已存在")
		}
	}
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
