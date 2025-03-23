package database

type GoShort struct {
	ID       uint64 `json:"id"       gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goshort  string `json:"goshort"  gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}
