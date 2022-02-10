package internal

import (
	"github.com/gin-gonic/gin"
	"test/interanl/core/entity"
	"test/interanl/core/interfaces"
)

func UpdateUser(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.UseCase)
	user := &entity.User{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(BadRequestHttpResponce(err))
	}
	id := c.Param("id")

	err = usecase.UpdateUser(id, user)
	if err != nil {
		c.JSON(BadRequestHttpResponce(err))
		return
	}
	c.JSON(SuccessHttpResponse(nil))
}
