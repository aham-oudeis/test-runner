package controllers

import (
	"test-runner/api/entities"
	"test-runner/api/utils/assertion_checker"
	"test-runner/api/utils/converter"

	"github.com/gin-gonic/gin"
)

func ValidateAssertions(c *gin.Context) {
	var assertionResponse entities.AssertionResponse

	if err := c.BindJSON(&assertionResponse); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request"})
		return
	}

	assertion := assertionResponse.Assertion
	response := assertionResponse.Response
	responseMap := converter.StructToMap(response)

	pass, actualValue, err := assertion_checker.IsAssertionValid(assertion, responseMap)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Assertion could not be validated because of:" + err.Error()})	
		return
	}
	
	//this is the part where I have to save the value into a database
	assertionResult := entities.AssertionResult{
		Pass: pass,
		ResponseId: response.Id,
		AssertionId: assertion.Id,
		Actual: actualValue,
	}

	c.JSON(200, assertionResult)
}
