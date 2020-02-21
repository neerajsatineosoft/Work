package main

import (
	"AddressApi/Common"
	"AddressApi/Model"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func initDb() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = gorm.Open(
		"postgres", os.Getenv("DB_URL"))

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected to db")
}
func AddUserAddress(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("error", err)
	}
	db.AutoMigrate(&Model.Address{})
	// bodyErr := json.NewDecoder(r.Body).Decode(&useraddress)
	// if err != nil {
	// 	Common.RespondWithError(w, http.StatusBadRequest, bodyErr.Error())
	// }
	rec := Model.Address{UserID: 2, City: "Aurangabad"}
	db.NewRecord(rec)
	db.Create(&rec)
}

func GetUserAddress(w http.ResponseWriter, r *http.Request) {
	allusersAddress := []Model.Address{}
	// db.AutoMigrate(&allusersAddress{})
	db.Find(&allusersAddress)
	fmt.Println("data of users is ", allusersAddress)
	Common.RespondWithJson(w, http.StatusOK, allusersAddress)
}
func main() {
	r := mux.NewRouter()
	initDb()
	r.HandleFunc("/GetAddress", GetUserAddress)
	r.HandleFunc("/PostAddress", AddUserAddress)
	http.ListenAndServe(":"+os.Getenv("HOST"), r)
	defer db.Close()
}
