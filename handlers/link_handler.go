package handlers

import (
	"links_available/services"

	"github.com/gin-gonic/gin"
)

type CheckLinksRequest struct {
	Links []string `json:"links" binding:"required"`
}
type GetLinksRequest struct {
	LinksList []int `json:"links_list" binding:"required"`
}

func CheckAvailableLinks(c *gin.Context) {
	var request CheckLinksRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат"})
		return
	}

	results := services.CheckLinks(request.Links)

	setID, err := services.SaveLinksSet(results)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка сохранения"})
		return
	}

	c.JSON(200, gin.H{
		"links":     results,
		"links_num": setID,
	})
}

func GetSavedLinks(c *gin.Context) {
	var request struct {
		LinksList []int `json:"links_list"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат"})
		return
	}

	pdfBytes := services.GeneratePDFReport(request.LinksList)

	c.Header("Content-Disposition", "attachment; filename=report.pdf")
	c.Data(200, "application/pdf", pdfBytes)
}
