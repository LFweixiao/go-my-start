package system

import (
	"archive/zip"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
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
// TODO zip的目录层级不对
// TODO 删除本地额 zip
func (f *FileServer) DownloadZip(c *gin.Context) error {

	// 这里用的绝对路径 读取到本地文件到流
	var packPath = global.PRO_CONFIG.System.ProjectPath + DIR_PATH
	dirs, err := ioutil.ReadDir(packPath)
	if err != nil {
		return err
	}

	fileName := "打包文件2"
	// 创建zip 打包文件
	localZip, err := os.Create(packPath + fileName + ".zip")
	if err != nil {
		return err
	}

	// 可操作的流
	zipWriter := zip.NewWriter(localZip)
	defer zipWriter.Close()

	// 读取目录中的文件打包
	for _, file := range dirs {
		err := compress(file, packPath, zipWriter, packPath)
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

// compres 根据文件 是dir 还是文件来 判断压缩
// file 传入文件
// dirPath 当前目录路径
// zzip zip读流
// zipPath 打包目录的路径
func compress(file os.FileInfo, dirPath string, zzip *zip.Writer, zipPath string) error {

	// 当前文件是目录时
	if file.IsDir() {
		// 拼接目录
		dirPathtwo := dirPath + file.Name() + "/"
		dir, err := ioutil.ReadDir(dirPathtwo)
		if err != nil {
			return err
		}
		// 读取目录下的文件
		for _, fi := range dir {
			f, err := os.Open(dirPathtwo + fi.Name())
			if err != nil {
				return err
			}
			defer f.Close()
			info, err := f.Stat()
			if err != nil {
				return err
			}
			err = compress(info, dirPathtwo, zzip, zipPath)
			if err != nil {
				return err
			}
		}
	} else {
		// 创建一个文 头部
		header, err := zip.FileInfoHeader(file)
		if err != nil {
			return err
		}
		// 绝对路径
		header.Name = dirPath + header.Name

		// 打开完整路径
		f, err := os.Open(header.Name)
		if err != nil {
			return err
		}
		// 路径改为打包目录的相对路径
		header.Name = header.Name[len(zipPath):]

		writer, err := zzip.CreateHeader(header) //这里创建文件时注意不要用完整路径 zip中会生产完整路径的目录
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, f)
		defer f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
