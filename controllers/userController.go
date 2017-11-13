package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/wjurmy/havaard/common"
	"github.com/wjurmy/havaard/data"
	"github.com/wjurmy/havaard/models"
	//"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"log"
)

// HTTP for HTTP Post - "user/register
// Add a new User Document

func Register(w http.ResponseWriter, r *http.Request){
	log.Println("in regitration handler")
	var dataResource UserResource

	//Decode the incoming User json
	//byt, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(byt))
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil{
		common.DisplayAppError(
			w,
			err,
			"Invalid User",
			500,
		)
		return
	}
	log.Printf("%#v",dataResource)
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("users")
	// count, err = c.Find(bson.M{email: value}).Count()

	repo := &data.UserRepository{c}


	repo.CreateUser(user)

	// Clean-up the hashpassword to eliminite it from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurde",
			500,
		)
		return
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func Login(w http.ResponseWriter, r *http.Request)  {

	var dataResouce LoginResource

	err := json.NewDecoder(r.Body).Decode(&dataResouce)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login data",
			500,
		)
		return
	}
	loginModel := dataResouce.Data
	loginUser := models.User{
		Email:	loginModel.Email,
		Password: loginModel.Password,
	}

	context :=NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}


	//authenticate the login user
	if user, err := repo.Login(loginUser); err != nil{
		common.DisplayAppError(
			w,
			err,
			"Invalid login creadentials",
			401,
		)
		return
	}else {
		// if login is successful
		token,err := common.GenerateJWT(user, "member")
		if err != nil{
			common.DisplayAppError(
				w,
				err,
				"Error while generation the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")

			user.HashPassword = nil
			authUser := AuthUserModel{
			User: user,
			Token: token,
			}

			j, err := json.Marshal(AuthUserResource{Data: authUser})
			if err != nil {
				common.DisplayAppError(
					w,
					err,
					"An unexpected error has occurd.",
					500,
				)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}

// Fetch User Data


func GetUserByID(w http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	user, err := repo.GetUserById(id)
	if err != nil {
		if err == mgo.ErrNotFound{
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return

		}
	}
	if j, err := json.Marshal(user); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
