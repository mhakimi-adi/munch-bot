package model

type Item struct {
	ItemId      int    `json:"item_id"`
	ItemName    string `json:"item_name"`
	ItemDesc    string `json:"item_desc"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	IsDeleted   int    `json:"is_deleted"`
}

func GetItems() ([]Item, error) {
	return items, nil
}

var items = []Item{
	{
		ItemId:      1,
		ItemName:    "Oil",
		ItemDesc:    "Oil",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
	{
		ItemId:      2,
		ItemName:    "Banana",
		ItemDesc:    "Banana",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
	{
		ItemId:      3,
		ItemName:    "Orange",
		ItemDesc:    "Orange",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
}
