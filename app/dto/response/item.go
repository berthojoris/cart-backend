package response

type Item struct {
	ID              uint   `json:"id"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	Image           string `json:"image"`
	Type            string `json:"type"`
	Price           int    `json:"price"`
}
