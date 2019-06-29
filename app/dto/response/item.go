package response

type Item struct {
	ID int `json:"id"`
	ItemName string `json:"item_name"`
	ItemDescription string `json:"item_description"`
}