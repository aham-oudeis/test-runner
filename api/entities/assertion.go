package entities

// the struct properties have to be capitalized to be exported; otherwise, they are private
// json tags are used to map the json keys to the struct properties
type Assertion struct {
	Property string `json:"property"`
	Comparison string `json:"comparison"`
	Expected string `json:"expected"`
	Id int `json:"id"`
}

type Request struct {
	Id int `json:"id"`
	Method string `json:"method"`
	Url string `json:"url"`
	Headers map[string]interface{} `json:"headers"`
	Body any `json:"body"`
	Assertions []Assertion `json:"assertions"`
}

type Response struct {
	Request Request `json:"request"`
	Status string `json:"status"`
	Latency int `json:"latency"`
	Body any `json:"body"`
	Headers map[string]interface{} `json:"headers"`
	Id int `json:"id"`
}

type AssertionResult struct {
	ResponseId int `json:"responseId"`
	Pass bool `json:"pass"`	
	Actual any `json:"actual"`
	AssertionId int `json:"assertionId"`
}

type AssertionResponse struct {
	Assertion Assertion `json:"assertion"`
	Response Response `json:"response"`
}
