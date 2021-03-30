package model
type Area struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	Name string `json:"name"`
	Description string `json:"description"`
}