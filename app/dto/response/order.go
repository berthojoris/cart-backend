package response

type Order struct {
	ID int `json:"id"`
	Orderid string `json:"order_id"`
	TotalAmount string `json:"total_amount"`
	CreatedBy int `json:"created_by"`
	CreatedDate interface{} `json:"created_by"`
	UpdatedDate interface{} `json:"created_by"`
}