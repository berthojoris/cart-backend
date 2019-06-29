package response

type OrderDetail struct {
	ID int `json:"id"`
	DetailOrderid string `json:"detail_order_id"`
	ItemId int `json:"item_id"`
	Qty string `json:"qty"`
}