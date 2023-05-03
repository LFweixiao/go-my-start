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
// @Tags     SystemFileApi
// @Summary  文件保存到本地
// @accept    multipart/form-data
// @Produce   application/json
// @Param    file  formData  file
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /systemFileApi/uploadLocal [post]
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

// DownloadZip 下载zip打包 本地已有的所有文件
// https://www.runoob.com/http/http-content-type.html 响应头设置
func (s *SystemFileApi) DownloadZip(c *gin.Context) {
	err := fileServer.DownloadZip(c)
	if err != nil {
		global.PRO_LOG.Error("打包失败", zap.Error(err))
	}
}
