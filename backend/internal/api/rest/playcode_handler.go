package rest

import (
	"log/slog"
	"net/http"

	"backend/internal/gocode"

	"github.com/gin-gonic/gin"
)

type playCodeRequest struct {
	Code string `json:"code" binding:"required"`
}

type playcodeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

// playcode godoc
// @Summary Play code
// @Description This endpoint is used to play the given code
// @Tags code
// @Accept json
// @Produce json
// @Param request body playCodeRequest true "Code to play"
// @Success 200 {object} playcodeResponse
// @Router /playcode [post]
func (api *apiDetails) playcode(c *gin.Context) {
	req := &playCodeRequest{}
	err := c.BindJSON(req)
	if err != nil {
		slog.Debug("error in binding json", err)
		createErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(req)
	if err != nil {
		createErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := api.app.Play(
		&gocode.GoCode{
			Code: req.Code,
		},
	)
	if err != nil {
		slog.Debug("error in getting response from app", err)
		createErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, &playcodeResponse{
		Output: resp.Output,
		Error:  resp.Error,
	})
	c.Done()
}
