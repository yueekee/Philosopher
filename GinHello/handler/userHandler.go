package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// user相关的转发接口
func UserSave(context *gin.Context) {
	// 通过Param方法获取参数
	username := context.Param("name")
	context.String(http.StatusOK, "用户:"+username+"已经保存")
}