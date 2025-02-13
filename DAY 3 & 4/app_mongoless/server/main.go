package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          string     `json:"id"`
	Number      string     `json:"number"`        //pascal case (first word starts with small letter and remaining words start with capital)
	AirlineName string     `json:"airline_name"` //snake case(two words seperated by _)
	Source      string     `json:"source"`
	Destination string     `json:"destination"`
	Capacity    int        `json:"capacity"`
	Price       float32    `json:"price"`
}

func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "1001", Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
		{Id: "1002", Number: "AI 673", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 18000.0},
	}
	c.JSON(http.StatusOK, flights)

}

func readFlightsById(c *gin.Context) {
	id := c.Param("id")
	flights := []Flight{
		{Id: id, Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
	}
	c.JSON(http.StatusOK, flights)
}
func createflight(c *gin.Context) {
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error" + err.Error()})
		return
	}
	createdflight := []Flight{
		{Id: "1000", Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
	}
	c.JSON(http.StatusCreated,
		gin.H{"message": "Flight created successfully.", "flight": createdflight})
}

func updateflight(c *gin.Context) {
	id := c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error" + err.Error()})
		return
	}
	updatedflight := []Flight{
		{Id: id, Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
	}
	c.JSON(http.StatusOK,
		gin.H{"message": "Flight updated successfully.", "flight": updatedflight})
}

func main() {
	// router
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	// routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightsById)
	r.POST("/flights", createflight)
	r.PUT("/flights/:id", updateflight)

	// server (default port 8080) //r.Run(":8080")
	r.Run(":8080")

	// flight1 := Flight {Id: "1001", Number: "AI 653", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	// fmt.Println(flight1)
}
