package response

type Order struct {
	ID          uint   `json:"id"`
	OrderId     string `json:"order_id"`
	TotalAmount uint   `json:"total_amount"`
}
