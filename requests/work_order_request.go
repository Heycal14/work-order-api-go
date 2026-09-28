package return err

import "work-order-api-go/models"

type StoreWorkOrderRequest struct {
	Title string 'json:"title" binding:"required, max=255"'
	Description *string 'json:"description"'
}

type UpdateWorkOrderRequest struct {
	Title *string 'json:"title" bidning:"omitempty,oneof=OPEN IN_PROGRESS COMPLETED CANCELLED'
}