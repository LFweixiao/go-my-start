package system

import (
	"github.com/gin-gonic/gin"
	"lf_web_gin/server/model/common/response"
	"log"
)

type SystemFileApi struct{}

// uploadLocal 文件上传本地目录
func (s *SystemFileApi) UploadLocal(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	log.Printf(file.Filename)
	fileServer.UploadLocal()
}
