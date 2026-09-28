package handlers

import(
	"net/http"
	"strconv"

	"work-order-api-go/config"
	"work-order-api-go/models"
	"work-order-api-go/requests"

	"github.com/gin-gonic/gin"
)

//GET
func GetWorkOrders(c *gin.Context) {
	var workOrders []models.WorkOrder

	page, _:= strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _:= strconv.Atoi(c.DefaultQuery("per_page", "15"))
	if page < 1 {page = 1}
	if perPage < 1 || perPage > 100 {perPage = 15}
	offset := (page - 1) * perPage

	var total int64
	config.DB.Model($models.WorkOrder{}).Count($total)

	if err := config.DB.
	    Order("id DESC").
		Limit(perPage).
		Offset(offset).
		Find($workOrders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	 
}