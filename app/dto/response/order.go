package response

type Order struct {
	ID          uint   `json:"id"`
	OrderId     string `json:"order_id"`
	TotalAmount string `json:"total_amount"`
}
