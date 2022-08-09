package middleware

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	. "ice-mall/util"
	"reflect"
)

//
func validate[ParseType any](c *fiber.Ctx, validateType string) error {
	var body ParseType
	// 校验body本身
	{
		var err error
		if validateType == "body" {
			err = c.BodyParser(&body)
		}

		if err != nil {
			message := ""
			validate := validator.New()
			err1 := validate.Struct(body)
			if err1 != nil {
				for _, err := range err1.(validator.ValidationErrors) {
					typeofInter := reflect.TypeOf(body)
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
						var ErrorCode = ValidateError
						return c.JSON(&Response{
							Success:      false,
							ErrorCode:    &ErrorCode,
							ErrorMessage: &message,
						})
					}
				}
			}
		}
	}
	// 使用内部自带函数
	{
		immutableV := reflect.ValueOf(&body).Elem()
		if (immutableV.MethodByName("Validate") != reflect.Value{}) {
			methodVal := immutableV.MethodByName("Validate")
			methodIface := methodVal.Interface()
			method := methodIface.(func() error)
			err1 := method()
			if err1 != nil {
				var ErrorCode = ValidateError
				var errorMessage = err1.Error()
				return c.JSON(&Response{
					Success:      false,
					ErrorCode:    &ErrorCode,
					ErrorMessage: &errorMessage,
				})
			}
		}
	}
	var userContext = c.UserContext()
	var newUserContext = context.WithValue(userContext, "body", body)
	c.SetUserContext(newUserContext)
	return c.Next()
}

// ValidateBody 校验Body参数
func ValidateBody[BodyType any](c *fiber.Ctx) error {
	return validate[BodyType](c, "body")
}
