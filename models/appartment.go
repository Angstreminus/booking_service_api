package models

type Appartment struct {
	ID       int                     `db:"id" json:"id"`
	BookID   int                     `db:"bookID" json:"bookID"`
	Type     string                  `db:"type" json:"type"`
	Capacity int                     `db:"capacity" json:"capacity"`
	Service  *AppartmentSeriviceList `db:"-" json:"service,omitempty"`
}
