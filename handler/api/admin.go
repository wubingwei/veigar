package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wubingwei/veigar/models"
	"net/http"
	"strconv"
)

func Admin(c *gin.Context) {
	idStr := c.Param("id")
	adminID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    4,
			"message": "id 不合法",
			"error":   "bad request",
		})
		return
	}
	admin, err := models.AdministratorStore.Find(adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    5,
			"message": "请重试一次",
			"error":   "server 查找失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    admin.ID,
		"name":  admin.Name,
		"email": admin.Email,
	})
	return

}
