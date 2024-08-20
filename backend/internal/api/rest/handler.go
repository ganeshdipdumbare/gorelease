package rest

import (
	_ "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var validate *validator.Validate

type errorRespose struct {
	ErrorMessage string `json:"errorMessage"`
}

func createErrorResponse(c *gin.Context, code int, message string) {
	c.IndentedJSON(code, &errorRespose{
		ErrorMessage: message,
	})
}

func (api *apiDetails) setupRouter() *gin.Engine {
	validate = validator.New()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Access-Control-Allow-Origin")
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// add middleware to add headers to the response
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Max-Age", "3600")
		c.Next()
	})
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	apiV1.POST("/playcode", api.playcode)
	apiV1.GET("/health", api.health)
	return r
}
