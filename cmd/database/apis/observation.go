package apis

import (
	"log"
	"strconv"
	"github.com/umeat/go-gnss/cmd/database/daos"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetObservation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if obs, err := daos.GetObservation(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, obs)
	}
}
