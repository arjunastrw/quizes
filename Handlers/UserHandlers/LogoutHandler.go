package UserHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LogoutHandler adalah fungsi untuk menangani permintaan logout pengguna.
func LogoutHandler(c *gin.Context) {
	c.SetCookie("jwtToken", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfuly"})
}
