package Model

// type User struct {
// 	SeqNo    int    `json:"seqno"`
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Birthday string `json:"birthday"`
// }
type Address struct {
	AddressID uint   `gorm:"primary_key"`
	City      string `json:"city"`
	UserID    uint
}
