package apis

import (
	"github.com/umeat/go-gnss/cmd/database/daos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetObservation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := daos.GetObservation(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
