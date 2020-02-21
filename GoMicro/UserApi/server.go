package main

import (
	"UserApi/Common"
	"UserApi/Model"
	"encoding/json"
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

func AddUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("error", err)
	}
	db.AutoMigrate(&Model.User{})
	allusers := Model.User{Name: "prashant", Age: 27, Birthday: "09sep"}
	db.NewRecord(allusers)
	db.Create(&allusers)
	Common.RespondWithJson(w, http.StatusOK, allusers)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	allusers := []Model.User{}
	db.Find(&allusers)
	fmt.Println("data of users is ", allusers)
	res, err := MakeRequest()
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	var userInfo []interface{}
	fmt.Println("after called to make request data res and allusers", res, allusers)
	for i := 0; i < len(allusers); i++ {
		userInfo = append(userInfo, allusers[i])
		for j := 0; j < len(res); j++ {
			if allusers[i].UserID == res[j].UserID {
				fmt.Println("from j loop ", allusers, res[j])
				userInfo = append(userInfo, res[j].City)
			}

		}
	}
	Common.RespondWithJson(w, http.StatusOK, userInfo)
}

func MakeRequest() ([]Model.Address, error) {
	r, err := http.Get("http://localhost:8082/GetAddress")
	if err != nil {
		return nil, err
		fmt.Println("error in url section")
	}
	defer r.Body.Close()
	args := []Model.Address{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		return nil, bodyErr
	}
	fmt.Println("after unmarshals", args)
	for i := 0; i < len(args); i++ {

	}
	return args, nil
}
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
func main() {
	r := mux.NewRouter()
	initDb()
	r.HandleFunc("/Getusers", GetUsers)
	r.HandleFunc("/AddUser", AddUser)
	http.ListenAndServe(":"+os.Getenv("HOST"), r)
	defer db.Close()
}
