package response

type OrderDetail struct {
	ID            uint `json:"id"`
	OrderId int  `json:"order_id"`
	ItemId        int  `json:"item_id"`
	Qty           int  `json:"qty"`
}
