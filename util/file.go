package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"ice-mall/config"
	"ice-mall/schema/annex"
	"os"
	"path"
	"strconv"
	"time"
)

// DirExists 文件夹是否存在
func DirExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// GetFileExt 获取文件后缀
func GetFileExt(filePath string) string {
	fileNameWithSuffix := path.Base(filePath)
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	return fileType
}

// GenerateFileName 生成文件名
func GenerateFileName(ext string) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return strconv.FormatInt(time.Now().Unix(), 10) + ext
	}
	return strconv.FormatInt(node.Generate().Int64(), 10) + ext
}

// CovertTypeFromFile 转换文件类型
func CovertTypeFromFile(fileType string) int {
	fmt.Println(fileType)
	for k, v := range config.FileMaps {
		for _, value := range v {
			if value == fileType {
				return k
			}
		}
	}
	return annex.UnknownFileType
}
