package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"test-runner/api/entities"
	"time"

	"github.com/gin-gonic/gin"
)

func ExecuteRequests(c *gin.Context) {
	var listOfRequests []entities.Request

	if err := c.BindJSON(&listOfRequests); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var listOfResponses []entities.Response

	client := http.DefaultClient
	for _, request := range listOfRequests {
		res, err := sendRequest(request, client)

		listOfResponses = append(listOfResponses, *res)

		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Request could not be sent because of:" + err.Error()})
			return
		}
	}

	c.JSON(200, listOfResponses)
}

func sendRequest(request entities.Request, client *http.Client) (*entities.Response, error) {
		body := request.Body;

		bodyString, err := json.Marshal(body) 
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(bodyString))
		if err != nil {
			return nil, err
		}

		timeBeforeRequest := time.Now()

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		timeAfterRequest := time.Now()

		latency := timeAfterRequest.Sub(timeBeforeRequest)

		response := entities.Response{
			Request: request,
			Status: strings.Split(res.Status, " ")[0],
			Latency: int(latency.Milliseconds()),
			Body: res.Body,
			//Headers: converter.StructToMap(res.Header),
		}

		return &response, nil
}