package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	v1 "lf_web_gin/server/api/v1"
	"lf_web_gin/server/global"
	"path"
	"runtime"
)

type FileApiRouter struct{}

func (f *FileApiRouter) InitFileApiRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	fileRouterApi := v1.ApiGroupApp.System.SystemFileApi

	//获取当前文件的路径
	_, fPath, _, ok := runtime.Caller(0) // /Applications/0utils/note/goProject/goMylfweixiao/go-my-start/server/router/system/sys_file_router.go
	if !ok {
		errors.New("获取本地文件路径失败！")
		return
	}
	//此方式是获取上级目录
	root := path.Dir(path.Dir(path.Dir(path.Dir(fPath))))

	//自己拼接跟目录 感觉不是很靠谱 不确定项目名 和文件名一直
	/*	index := strings.Index(fPath, "go-my-start")
		s := fPath[0:index+len("go-my-start")] + "/"*/

	global.PRO_CONFIG.System.ProjectPath = root + "/"

	{
		fileRouter.POST("uploadLocal", fileRouterApi.UploadLocal)         // 文件上传本地
		fileRouter.POST("removeLoaclFile", fileRouterApi.RemoveLocalFile) //删除本地文件
		fileRouter.POST("dwonloadZip", fileRouterApi.DownloadZip)         // 打包本地已有文件
	}

}
