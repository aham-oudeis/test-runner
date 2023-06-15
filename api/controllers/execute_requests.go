package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
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

	listOfResponses, err := MakeRequests(listOfRequests)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Request could not be sent because of:" + err.Error()})
	}

	c.JSON(200, listOfResponses)
}

func MakeRequests(listOfRequests []entities.Request) ([]entities.Response, error) {
	var listOfResponses []entities.Response
	
	orderedRequests := orderRequest(listOfRequests)

	client := http.DefaultClient

	for _, listOfRequest := range orderedRequests {
		//parse the request, this part is yet to be implemented: parsing happens every loop; once the parsing is done, 
		// all the parsed requests are made using goroutines 
		responses := make(chan entities.Response, len(listOfRequest))

		var wg sync.WaitGroup
		wg.Add(len(listOfRequest))
		for _, request := range listOfRequest {
			go makeOneRequest(request, client, responses, &wg)
		}
		wg.Wait()

		gatherResponses(responses, &listOfResponses) 

	}

	return listOfResponses, nil
}

func gatherResponses(responses chan entities.Response, listOfResponses *[]entities.Response) {
	for response := range responses {
		*listOfResponses = append(*listOfResponses, response)
	}
}

func orderRequest(listOfRequests []entities.Request) [][]entities.Request {
	orderedRequests := make([][]entities.Request, maxOrderNumber(listOfRequests))

	for _, request := range listOfRequests {
		orderNumber := request.OrderId
		orderedRequests[orderNumber] = append(orderedRequests[orderNumber], request)
	}

	fmt.Println(orderedRequests)

	return orderedRequests
}

func maxOrderNumber(listOfRequests []entities.Request) int {
	maxVal := 0

	for _, req := range listOfRequests {
		if req.OrderId > maxVal {
			maxVal = req.OrderId
		}
	}

	return maxVal
}

func makeOneRequest(request entities.Request, client *http.Client, responses chan entities.Response, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := sendOneRequest(request, client)
	if err != nil {
		responses <- entities.Response{
			Request: request,
			Status:  "Error",
			Latency: 0,
			Body:    nil,
			Headers: nil,
		}

		return
	}

	responses <- *res
}

func sendOneRequest(request entities.Request, client *http.Client) (*entities.Response, error) {
	body := request.Body

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
		Status:  strings.Split(res.Status, " ")[0],
		Latency: int(latency.Milliseconds()),
		Body:    res.Body,
		//Headers: converter.StructToMap(res.Header),
	}

	return &response, nil
}
