package dto

type CreateMenuDto struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description" binding:"required"`
    Price       float64 `json:"price" binding:"required,gt=0"`
    Qty         int32   `json:"qty" binding:"required,gte=0"`
}


