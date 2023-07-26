package handler

import (
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Openings
// @Description List openings
// @Tags Openings
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
