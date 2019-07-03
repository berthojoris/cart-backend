package response

type Order struct {
	ID          uint `json:"id"`
	OrderId     int  `json:"order_id"`
	TotalAmount uint `json:"total_amount"`
}
