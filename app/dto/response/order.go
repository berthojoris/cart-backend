package response

type Order struct {
	ID          uint `json:"id"`
	OrderId     uint `json:"order_id"`
	TotalAmount uint `json:"total_amount"`
}
