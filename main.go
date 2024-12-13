package main

import (
	"log"
	"net/http"

	"com.sunfish.ivfsuccesscalculator/controllers"
	"com.sunfish.ivfsuccesscalculator/models"
	"github.com/gin-gonic/gin"
)

func main() {
	formulas, err := models.LoadFormulasFromCSV("ivf_success_formulas.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize controller
	controller := controllers.NewIVFController(formulas)

	r := gin.Default()
	// CORS middleware for HTMX requests
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, HX-Request, HX-Current-URL")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Add root endpoint to serve the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/calculate", controller.CalculateSuccess)

	r.Run(":8081")

}
