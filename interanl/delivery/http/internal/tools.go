package internal

import (
	"net/http"
	"test/interanl/core/entity"
)

func SuccessHttpResponse(data interface{}) (int, entity.Dict) {
	if data == nil {
		return http.StatusOK,
			entity.Dict{"success": true}
	}

	return http.StatusOK,
		entity.Dict{
			"success": true,
			"data":    data,
		}
}

func BadRequestHttpResponce(err error) (int, entity.Dict) {

	return http.StatusBadRequest,
		entity.Dict{
			"success": false,
			"data":    err.Error(),
		}
}
