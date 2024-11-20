package main

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        string        `json:"total"`
	Items        []ReceiptItem `json:"items"`
}

type ProcessReceiptResponse struct {
	Id string `json:"id"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}
