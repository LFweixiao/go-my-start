package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileServer struct {
	Path string
}

// 本地文件临时相对路径
const DIR_PATH = "resources/file/"

// UploadLocal 本地文件传输
func (f *FileServer) UploadLocal(file *multipart.FileHeader, c *gin.Context) (*model.SysFile, error) {

	projectPath, fileV := timeFirUUIDName(file.Filename)

	err := c.SaveUploadedFile(file, projectPath)
	if err != nil {
		return nil, errors.New("文件保存失败")
	}
	return fileV, nil
}

// timeFirAndFileName 生成时间+UUID的name
func timeFirUUIDName(fileName string) (string, *model.SysFile) {
	date := time.Now()
	index := strings.Index(fileName, ".")
	fileV := &model.SysFile{UUIDName: uuid.NewV1().String(), Postfix: fileName[index:], Name: fileName}
	// 月份需要格式化数字
	return global.PRO_CONFIG.System.ProjectPath + DIR_PATH + strconv.Itoa(date.Year()) + "/" + date.Format("01") + "/" + fileV.UUIDName + fileV.Postfix, fileV
}

// RemoveLocalFile 删除本地文件
func (f *FileServer) RemoveLocalFile(file *model.SysFile, c *gin.Context) (*model.SysFile, error) {

	path := file.Year + "/" + file.Month + "/" + file.UUIDName + file.Postfix
	path = global.PRO_CONFIG.System.ProjectPath + DIR_PATH + path
	err := os.Remove(path)
	if err != nil {
		return file, errors.New("文件不存在、或已经删除")
	}
	return nil, nil
}
