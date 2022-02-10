package internal

import (
	"github.com/gin-gonic/gin"
	"test/interanl/core/interfaces"
)

func GetUser(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.UseCase)
	id := c.Param("id")
	user, err := usecase.GetUser(id)
	if err != nil {
		c.JSON(BadRequestHttpResponce(err))
		return
	}
	c.JSON(SuccessHttpResponse(user))
}
