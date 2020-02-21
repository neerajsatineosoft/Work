package Model

type User struct {
	UserID   uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	// Address  Address `json:"address"`
	// Address []Address `gorm:"foreignkey:AddressID"`
}
type Address struct {
	AddressID uint   `json:"user_id`
	City      string `json:"city"`
	UserID    uint   `json:"user_id`
}

// type Result struct {
// 	UserID   int
// 	Name     string
// 	Age      int
// 	Birthday string
// 	Address  []Address
// }
