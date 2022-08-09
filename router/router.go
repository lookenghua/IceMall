package router

import (
	"fmt"
	"github.com/goccy/go-json"
	. "ice-mall/common/dto"
	"ice-mall/config"
	v1 "ice-mall/router/api/v1"
	. "ice-mall/router/guard"
	. "ice-mall/router/middleware"
	"ice-mall/util"
	"os"
)
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// InitRouter 初始化路由
func InitRouter() *fiber.App {
	// 自定义json转换器,提高效率
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   50 * 1024 * 1024, // 50M限制
	})

	// 全局中间件
	{
		// 创建静态资源托管
		{
			app.Static("/file", config.StoragePath)
			if !util.DirExists(config.StoragePath) {
				fmt.Println("创建uploads文件夹")
				err := os.Mkdir(config.StoragePath, os.ModePerm)
				if err != nil {
					fmt.Println("创建uploads文件夹失败", err.Error())
				} else {
					fmt.Println("创建uploads文件夹成功")
				}
			}
		}

		// 异常中间件
		app.Use(recover.New())
		// 压缩中间件
		app.Use(compress.New(compress.Config{
			Level: compress.LevelBestSpeed, // 1
		}))
		// 跨域中间件
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
		// 日志中间件
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
		// RequestID中间件
		app.Use(requestid.New(requestid.Config{
			Header: "X-Custom-Header",
			Generator: func() string {
				return "static-id"
			},
		}))
		// 日志记录中间件
		app.Use(LoggerMiddleware)
	}
	// 全局路由
	{
		// 展示监控
		app.Get("/", monitor.New())
	}

	// 路由
	{
		api := app.Group("/api")
		apiV1 := api.Group("/v1")
		// 登录注册
		{
			// 管理员登录
			apiV1.Post("/admin/token", ValidateBody[LoginAdminDto], v1.LoginAdmin)
			// 创建用户
			apiV1.Post("/user", ValidateBody[CreateUserDto], v1.CreateUser)
			// 用户登录
			apiV1.Post("/user/token", ValidateBody[LoginUserDto], v1.LoginUser)
			// 获取用户登录信息
			apiV1.Get("/user", UserGuard, v1.GetCurrentUserInfo)
		}
		// 商品
		{
			// 创建商品
			apiV1.Post("/product", UserGuard, ValidateBody[CreateProductDto], v1.CreateProduct)
		}
		// 工具
		{
			apiV1.Post("/upload", UserGuard, ValidateBody[UploadFileDto], v1.UploadFile)
		}
	}

	return app
}
