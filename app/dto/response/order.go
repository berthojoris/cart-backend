package response

type Order struct {
	ID          uint `json:"id"`
	TotalAmount uint `json:"total_amount"`
}
