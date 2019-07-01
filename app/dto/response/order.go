package response

type Order struct {
	ID          uint        `json:"id"`
	OrderId     string      `json:"order_id"`
	TotalAmount string      `json:"total_amount"`
	CreatedBy   int         `json:"created_by"`
	CreatedDate interface{} `json:"created_date"`
	UpdatedDate interface{} `json:"updated_date"`
}
