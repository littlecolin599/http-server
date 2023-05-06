package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (DefaultController) Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}
