package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model"
	"lf_web_gin/server/model/common/response"
)

type SystemFileApi struct{}

// uploadLocal 文件上传本地目录
func (s *SystemFileApi) UploadLocal(c *gin.Context) {

	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileV, err2 := fileServer.UploadLocal(file, c)
	if err2 != nil {
		global.PRO_LOG.Error("本地文件保存失败", zap.Error(err))
		response.FailWithMessage("本地文件保存失败", c)
	}
	response.OkWithData(fileV.UUIDName+fileV.Postfix, c)
}

// Remove 删除本地文件
func (s *SystemFileApi) RemoveLocalFile(c *gin.Context) {
	var fileVO *model.SysFile
	err := c.ShouldBindJSON(&fileVO)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = fileServer.RemoveLocalFile(fileVO, c)
	if err != nil {
		global.PRO_LOG.Error("本地文件删除失败", zap.Error(err))
		response.FailWithMessage("本地文件删除失败", c)
	}
	response.Ok(c)
}
