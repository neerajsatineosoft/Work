package Controller

/*
import (
	"UserApi/Model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB_URL = os.Getenv("DB_URL")

func AddUser() {
	db, err := gorm.Open("postgres", DB_URL)
	if err != nil {
		fmt.Println("error", err)
	}
	// defer db.Close()
	user := Model.User{Name: "Jinzhu", Age: 18, Birthday: "092"}
	db.Create(&user)
}
func GetAllUsers() (users []Model.User) {
	allusers := []Model.User{}
	db.Find(&allusers)
	fmt.Println("data of users is ", allusers)
	return allusers
}
*/
