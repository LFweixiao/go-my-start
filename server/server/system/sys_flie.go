package system

import (
	"archive/zip"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
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

// DownloadZip 下载Zip文件
// 下载指定路径文件，并且会在本地生产 zip包
// TODO ZIP文件怎么通过gin 传回前端
func (f *FileServer) DownloadZip(c *gin.Context) error {

	// 这里用的绝对路径 读取到本地文件到流
	var packPath = global.PRO_CONFIG.System.ProjectPath + DIR_PATH
	dirs, err := ioutil.ReadDir(packPath)
	if err != nil {
		return err
	}

	fileName := "打包文件"
	// 创建zip 打包文件
	localZip, err := os.Create(packPath + fileName + ".zip")
	if err != nil {
		return err
	}

	// 可操作的流
	zipWriter := zip.NewWriter(localZip)
	defer zipWriter.Close()

	// 读取本地文件 写入zip
	for _, file := range dirs {
		// zip包中创建同名文件
		readFile, err := zipWriter.Create(file.Name())
		fileBytes, err := ioutil.ReadFile(packPath + file.Name())
		if err != nil {
			return err
		}
		// 写入文件内容
		_, err = readFile.Write(fileBytes)
		if err != nil {
			return err
		}
	}

	// 设置响应头
	c.Header("Content-Dispostition", "attachment;filename="+fileName+".zip") // 文件名
	c.Header("Content-Type", "application/zip")                              // 文件传输类型 .zip
	//c.Writer.Write()
	return nil
}
