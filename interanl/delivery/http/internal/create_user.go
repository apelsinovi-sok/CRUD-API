package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"test/interanl/core/entity"
	"test/interanl/core/interfaces"
	"time"
)

func CreateUser(c *gin.Context) {
	usecase := *c.MustGet("usecase").(*interfaces.UseCase)
	user := &entity.User{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(BadRequestHttpResponce(err))
		return
	}
	user.ID = uuid.New()
	user.Created = time.Now()

	err = usecase.CreateUser(user)
	if err != nil {
		c.JSON(BadRequestHttpResponce(err))
		return
	}
	c.JSON(SuccessHttpResponse(nil))
}
