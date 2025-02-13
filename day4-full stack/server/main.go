package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct{
	Id 				string 			`json:"id"`			  // changing the key to required format
	Number 			string			`json:"number"`		  // Pascal case - letter capitalised after every word
	AirlineName		string			`json:"airline_name"` //snake case - joined by underscore
	Source 			string			`json:"source"`
	Destination 	string			`json:"destination"`
	Capacity 		int				`json:"capacity"`
	Price 			float32			`json:"price"`
}

func readAllFlights( c *gin.Context) {
	flights := [] Flight {
		{Id: "1001", Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
		{Id: "1002", Number: "AI 673", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 18000.0},
	}
	c.JSON(http.StatusOK, flights)

}

func readFlightsById (c *gin.Context) {
	id := c.Param("id")
	flight := Flight {Id: id, Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0 }

	c.JSON(http.StatusOK, flight)
}

func createFlight( c *gin.Context) {
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"Error" : "Server Error." +err.Error()})
		return
	}

	createdFlight :=Flight {
		Id: "1005", Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusCreated, 
		gin.H{"message" : "Flight created Successfully", "flight": createdFlight})	
}

func updateFlight( c *gin.Context) {
	id := c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"Error" : "Server Error." +err.Error()})
		return
	}

	updatedFlight :=Flight {
		Id: id , Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusOK, 
		gin.H{"message" : "Flight updated Successfully", "flight": updatedFlight})	
}

func deleteFlight( c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, 
		gin.H{"message" : "Flight deleted Successfully"})	
}

func main() {
	// router
	r := gin.Default()
	// routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightsById)
	// server (default port 8080) //r.Run(":8080")
	
	r.POST("/flights", createFlight)	
	r.PUT("/flights/:id", updateFlight)	
	r.DELETE("/flights/:id", deleteFlight)	
	r.Run(":8080") // specifying the port
	//fmt.Println(flights)
}