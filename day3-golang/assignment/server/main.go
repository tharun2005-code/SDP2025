package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pizza struct{
	Id string
	Name string
	Size string
	Category string
	Price float32
}

func readAllPizzas( c *gin.Context) {
	pizzas := [] Pizza {
		{Id: "1001", Name: "Margherita", Size: "25 cm", Category: "Italian",  Price: 350.0},
		{Id: "1002", Name: "Farmhouse", Size: "60 cm", Category: "Indian",  Price: 200.0},
		{Id: "1003", Name: "Pepperoni", Size: "60 cm", Category: "Chinese",  Price: 250.0},
	}
	c.JSON(http.StatusOK, pizzas)

}

func readPizzasById (c *gin.Context) {
	id := c.Param("id")
	pizza := Pizza {Id: id, Name: "Margherita", Size: "25 cm", Category: "Italian",  Price: 350.0}

	c.JSON(http.StatusOK, pizza)
}

func createPizza( c *gin.Context) {
	var jbodyPizza Pizza
	err := c.BindJSON(&jbodyPizza)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"Error" : "Server Error." +err.Error()})
		return
	}

	createdPizza :=Pizza {
		Id: "1003", Name: "Pepperoni", Size: "60 cm", Category: "Chinese",  Price: 250.0}
	c.JSON(http.StatusCreated, 
		gin.H{"message" : "Pizza created Successfully", "pizza": createdPizza})	
}

func updatePizza( c *gin.Context) {
	id := c.Param("id")
	var jbodyPizza Pizza
	err := c.BindJSON(&jbodyPizza)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"Error" : "Server Error." +err.Error()})
		return
	}

	updatedPizza :=Pizza {
		Id: id, Name: "Margherita", Size: "25 cm", Category: "Italian",  Price: 350.0}
	c.JSON(http.StatusOK, 
		gin.H{"message" : "Pizza updated Successfully", "pizza": updatedPizza})	
}

func deletePizza( c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, 
		gin.H{"message" : "Pizza deleted Successfully"})	
}

func main() {
	// router
	r := gin.Default()
	// routes | API Endpoints
	r.GET("/pizzas", readAllPizzas)
	r.GET("/pizzas/:id", readPizzasById)
	// server (default port 8080) //r.Run(":8080")
	
	r.POST("/pizzas", createPizza)	
	r.PUT("/pizzas/:id", updatePizza)	
	r.DELETE("/pizzas/:id", deletePizza)	
	r.Run(":8080")
	//fmt.Println(pizzas)
}