package handler

import (
	"github.com/Cloude77/mailing-system/api-service/internal/model"
	"github.com/Cloude77/mailing-system/api-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Annotation for Swagger

// @Summary Create a new mailing
// @Description Creates a new mailing task
// @Tags mailings
// @Accept json
// @Produce json
// @Param mailing body model.MailingRequest true "Mailing data"
// @Success 201 {object} model.Mailing
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/mailing [post]
func CreateMailing(s *service.MailingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.MailingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			logrus.Errorf("Invalid request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mailing, err := s.CreateMailing(c.Request.Context(), &req)
		if err != nil {
			logrus.Errorf("Failed to create mailing: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, mailing)
	}
}
