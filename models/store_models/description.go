package models

type StoreData struct {
	Categories []StoreCategory `json:"categories,omitempty"`
	Products   []StoreProduct  `json:"products,omitempty"`
}
