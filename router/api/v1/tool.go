package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	. "ice-mall/common/dto"
	"ice-mall/config"
	"ice-mall/service"
	. "ice-mall/util"
	"os"
	"time"
)

// UploadFile 上传文件
func UploadFile(c *fiber.Ctx) error {
	apiUtil := ApiUtil{
		Ctx: c,
	}
	data := GetBody[UploadFileDto](c)
	annexNum := data.AnnexNum
	file, err := c.FormFile("file")
	userInfo := apiUtil.GetUserInfo()
	userId := userInfo.ID
	if err != nil {
		return apiUtil.Fail(ValidateError, err.Error())
	}
	// 保存文件
	{
		filename := file.Filename
		ext := GetFileExt(filename)
		storageFileName := GenerateFileName(ext)
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		subDir := fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
		relativeSubDir := fmt.Sprintf(".%s/%s", config.StoragePath, subDir)
		if !DirExists(relativeSubDir) {
			err := os.Mkdir(relativeSubDir, os.ModePerm)
			if err != nil {
				fmt.Println("创建文件夹失败", err.Error())
				return apiUtil.Fail(UploadError, err.Error())
			}
		}
		fileUrl := fmt.Sprintf("%s/%s/%s", config.StoragePath, subDir, storageFileName)
		savePath := fmt.Sprintf(".%s", fileUrl)
		// 保存文件
		if err := c.SaveFile(file, savePath); err != nil {
			return err
		}
		result, saveErr := service.SaveFileToAnnex(annexNum, file, storageFileName, fileUrl, userId)
		if saveErr != nil {
			return apiUtil.Fail(SaveDataError, saveErr.Error())
		}
		return apiUtil.Success(result)
	}

}
