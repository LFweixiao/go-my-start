package system

import (
	"github.com/gin-gonic/gin"
	v1 "lf_web_gin/server/api/v1"
)

type FileApiRouter struct{}

func (f *FileApiRouter) InitFileApiRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	fileRouterApi := v1.ApiGroupApp.System.SystemFileApi

	{
		fileRouter.POST("uploadLocal", fileRouterApi.UploadLocal) // 文件上传本地
	}

}
