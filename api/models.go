package models

type ReceiptItem struct {
	ShortDescription str `json: "shortDescription validate: "required"`
	Price            str `json: "price" validate: "required"`
}

type Receipt struct {
	Retailer     str           `json: "retailer" validate: "required"`
	PurchaseDate str           `json: "purchaseDate validate: "required"`
	PurchaseTime str           `json: "purchaseTime validate: "required"`
	Total        str           `json: "total" validate: "required"`
	Items        []ReceiptItem `json: "items" validate: "required"`
}

type ProcessReceiptResponse struct {
	Id str `json: "id"`
}

type GetPointsResponse struct {
	Points int `json: "points"`
}
