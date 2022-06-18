package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Represents the expected JSON structure of the datasource. Customise to suit the data coming from the source API.
type PriceDataSource struct {
	Raw struct {
		Eth struct {
			Usd struct {
				Price float64 `json:"PRICE"`
			}
		} `json:"ETH"`
	} `json:"RAW"`
}

// Represents the JSON structure of the EA's data output. Customise the contents of this struct.
type DataOutput struct {
	Price float64 `json:"price"`
}

// Represents the JSON structure of the overall response. Should NOT need customisation.
type ExternalAdapterOutput struct {
	Data  *DataOutput `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type DataInput struct {
	SomeValue int `json:"someValue"`
}

// Main request handler of this external adapter. Customise to your heart's content.
func mainHandler(c *gin.Context) {
	defer requestsProcessed.Inc() // increases the metrics counter at the end of the request

	var dataInput DataInput
	dataInput.SomeValue = 10

	if c.BindJSON(&dataInput) != nil {
		c.JSON(406, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"someValue": dataInput.SomeValue})

}
