package models

type StoreSetting struct {
	Id        string `db:"id" json:"id"`
	PartnerId int    `db:"partner_id" json:"partner_id"`
	StoreId   int    `db:"store_id" json:"store_id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Currency  string `db:"currency" json:"currency"`
	Address   string `db:"address" json:"address"`
	SiretNo   string `db:"siret_no" json:"siret_no"`
	Logo      string `db:"logo" json:"logo"`
	DayRef    string `db:"day_reference" json:"day_reference"`
}
