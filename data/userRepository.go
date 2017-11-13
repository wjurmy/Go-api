package data

import (
	"github.com/wjurmy/havaard/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"


)

type UserRepository struct {
	C *mgo.Collection
}



func (r *UserRepository) CreateUser(user *models.User) error  {
	obj_id := bson.NewObjectId()
	user.Id = obj_id

	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.HashPassword = hpass
	user.CreatedOn = time.Now()

	//clear the incoming text password
	user.Password = ""

	err = r.C.Insert(user)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {

		return
	}
	// validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}


func (r *UserRepository) GetUserById(Id string) (user models.User, err error) {
	err = r.C.FindId(bson.ObjectIdHex(Id)).One(&user)
	return
}
