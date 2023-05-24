package models

type AppartmentSeriviceList struct {
	ID           int  `db:"id" json:"id"`
	AppartmentID int  `db:"appartmentID" json:"appartmentID"`
	Internet     bool `db:"internet" json:"internet"`
	Storage      bool `db:"storage" json:"storage"`
	Gym          bool `db:"gym" json:"gym"`
	Delivery     bool `db:"delivery" json:"delivery"`
	ClothCleaner bool `db:"clothcleaner" json:"clothcleaner"`
}
