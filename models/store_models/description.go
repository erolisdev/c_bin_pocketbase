package models

type Description struct {
	LanguageID       *int64  `json:"language_id,omitempty"`
	Name             *string `json:"name,omitempty"`
	DescriptionValue *string `json:"description,omitempty"`
}
