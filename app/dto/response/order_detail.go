package response

type OrderDetail struct {
	ID uint `json:"id"`
	DetailOrderid string `json:"detail_order_id"`
	ItemId int `json:"item_id"`
	Qty int `json:"qty"`
}