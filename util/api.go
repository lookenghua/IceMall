package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"ice-mall/schema"
	"reflect"
)

// 错误代码
const (
	ParseError              = iota + 1001 // 解析错误
	ValidateError                         // 校验错误
	UnknownError                          // 未知错误
	DataNotFoundError                     // 数据不存在
	DataFoundError                        // 数据已存在
	CreateError                           // 创建数据错误
	QUERYError                            // 查询数据错误
	PasswordNotMatchedError               // 密码不匹配
	CreateTokenError                      // 创建token失败
	SaveDataError                         // 保存数据失败
	IllegalToken                          // 非法token
	AuthorityNotRightError                // 权限不正确
	UploadError                           // 上传失败
)

// ApiUtil API工具
type ApiUtil struct {
	Ctx *fiber.Ctx
}

// Response 相应格式
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
	ErrorCode    *int        `json:"errorCode,omitempty"`
}

// ValidateBody 校验body
func (app *ApiUtil) ValidateBody(data interface{}) error {
	if err := app.Ctx.BodyParser(data); err != nil {
		if err != nil {
			fmt.Println(err)
			return fiber.NewError(ParseError, "解析失败")
		}

	}
	message := ""
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			typeofInter := reflect.TypeOf(data)
			if typeofInter.Kind() == reflect.Ptr {
				typeofInter = typeofInter.Elem()
			}
			if interType, ok := typeofInter.FieldByName(err.Field()); ok {
				jsonTag := interType.Tag.Get("json")
				var tag = err.Tag()
				switch tag {
				case "required":
					message = fmt.Sprintf("缺少%s参数", jsonTag)
				}
				return fiber.NewError(ValidateError, message)
			}
		}
	}
	return nil
}

// Success 成功返回
func (app *ApiUtil) Success(data interface{}) error {
	return app.Ctx.JSON(&Response{
		Success: true,
		Data:    data,
	})
}

// Fail 失败返回
func (app *ApiUtil) Fail(errorCode int, errorMessage string) error {
	return app.Ctx.JSON(&Response{
		Success:      false,
		ErrorCode:    &errorCode,
		ErrorMessage: &errorMessage,
	})
}

// SetUserInfo 设置用户信息
func (app *ApiUtil) SetUserInfo(user *schema.User) {
	app.Ctx.Locals("user", user)
}

// GetUserInfo 获取用户信息
func (app *ApiUtil) GetUserInfo() *schema.User {
	return app.Ctx.Locals("user").(*schema.User)
}

// Transform 数据转换
func (app *ApiUtil) Transform(from interface{}, to interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		fmt.Printf("转换类型失败")
		return
	}
}

// GetJsonBody 获取参数
func GetJsonBody[BodyType any](c *fiber.Ctx) BodyType {
	return c.UserContext().Value("jsonBody").(BodyType)
}
