package adminPanel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAdminPanel(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
