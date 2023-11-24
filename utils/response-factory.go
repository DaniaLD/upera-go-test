package utils

import (
	"encoding/json"
	"fmt"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func okConvertor(myInterface interface{}) (globalModel.OkModel, error) {
	em := globalModel.OkModel{}
	myMap, ok := myInterface.(map[string]interface{})
	if !ok {
		return em, fmt.Errorf("invalid input type")
	}

	if len(myMap) == 2 {
		var keys []string
		for key, _ := range myMap {
			keys = append(keys, key)
		}
		if keys[0] == "message" && keys[1] == "payload" {
			em.Message = myMap[keys[0]].(string)
			em.Payload = myMap[keys[1]]
			return em, nil
		}
		if keys[1] == "message" && keys[0] == "payload" {
			em.Message = myMap[keys[1]].(string)
			em.Payload = myMap[keys[0]]
			return em, nil
		}
	}

	em.Message = "ok"
	em.Payload = myMap
	return em, nil

}

func errorConvertor(myInterface interface{}) (globalModel.ErrorModel, error) {
	em := globalModel.ErrorModel{}
	myMap, ok := myInterface.(map[string]interface{})
	if !ok {
		return em, fmt.Errorf("invalid input type")
	}
	for key, value := range myMap {
		switch key {
		case "message":
			em.Message = value.(string)
		case "payload":
			em.Payload = value.(interface{})
		default:
			return em, fmt.Errorf("unknown key: %s", key)
		}
	}
	return em, nil
}

func okFormatter(payload interface{}) globalModel.ResponseModel {
	okModel, _ := okConvertor(payload)
	response := globalModel.ResponseModel{
		Status:  http.StatusOK,
		Message: okModel.Message,
		Payload: okModel.Payload,
	}
	return response
}

func errorFormatter(statusCode int, payload interface{}) globalModel.ResponseModel {
	errModel, _ := errorConvertor(payload)
	response := globalModel.ResponseModel{
		Status:  statusCode,
		Message: errModel.Message,
		Payload: errModel.Payload,
	}
	return response
}

func ResponseFormatter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return err
		}
		responseBody := c.Response().Body()

		var responseInterface interface{}
		err = json.Unmarshal(responseBody, &responseInterface)
		if err != nil {
			return err
		}
		statusCode := c.Response().StatusCode()
		var resp interface{}
		if statusCode == http.StatusOK {
			resp = okFormatter(responseInterface)
		} else {
			resp = errorFormatter(statusCode, responseInterface)
		}

		c.Status(statusCode).JSON(resp)
		return nil
	}
}
