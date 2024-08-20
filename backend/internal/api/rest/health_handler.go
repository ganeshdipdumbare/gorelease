package rest

import "github.com/gin-gonic/gin"

// health godoc
// @Summary Health check endpoint
// @Description This endpoint is used to check the health of the server
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} string "ok"
// @Router /health [get]
func (api *apiDetails) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
