package Api

import (
	"UserApi/Model"
	"fmt"
	"net/http"
)

/*
import (
	"UserApi/Common"
	"UserApi/Controller"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := Controller.GetAllUsers()
	Common.RespondWithJson(w, http.StatusOK, users)
}
*/
func GetUsers(w http.ResponseWriter, r *http.Request) {
	allusers := []Model.User{}
	db.Find(&allusers)
	fmt.Println("data of users is ", allusers)
}
