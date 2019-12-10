package adminPanel

import "github.com/gin-gonic/gin"

func InitAdmin(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
}
