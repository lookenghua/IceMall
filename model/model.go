package model

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"ice-mall/ent"
	"ice-mall/ent/migrate"
	"ice-mall/ent/systemconfig"
	"log"
)

var client *ent.Client

// InitDatabase 初始化数据库连接
func InitDatabase() {
	var err error
	client, err = ent.Open("mysql", "root:123456@tcp(localhost:3306)/ice_mall?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	SeedDatabase()
}

// SeedDatabase 初始化数据库数据
func SeedDatabase() {
	fmt.Printf("===开始初始化数据===\n")
	// 设置验证码发送间隔
	{
		n, err := client.SystemConfig.Query().Where(systemconfig.RuleEQ("SEND_CODE_TIME")).Count(context.Background())
		if err != nil && n == 0 {
			_, e := client.SystemConfig.Create().SetRule("SEND_CODE_TIME").SetValue("5").SetRemark("发送验证码间隔").Save(context.Background())
			if e != nil {
				fmt.Printf("%v\n", e)
				return
			}
		}
	}

	fmt.Println("===初始化数据成功===")
}
