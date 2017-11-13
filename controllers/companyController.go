package controllers

import (
	"net/http"

	"encoding/json"
	"github.com/wjurmy/havaard/common"
	"github.com/wjurmy/havaard/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"net/smtp"
	"fmt"
)

func RegisterCompany(w http.ResponseWriter, req *http.Request){
	var dataResource CompanyResource

	// Decoding the incoming company
	err := json.NewDecoder(req.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid company data",
			500,
		)
		return
	}

			company := &dataResource.Data
			context := NewContext()
			defer context.Close()

			c := context.DbCollection("companies")
			repo := &data.CompanyRepository{c}

			// insert a company document
			repo.CreateCompany(company)
			if j, err := json.Marshal(CompanyResource{Data: *company}); err != nil {
				common.DisplayAppError(
					w,
					err,
					"An unexpected error has occurred",
					500,
				)
				return

			} else {

				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusCreated)

				w.Write(j)
			}

}

func UpdateCompany(w http.ResponseWriter, req *http.Request){
	// Get Id from the incoming url
	log.Print("Entered update...")
	vars := mux.Vars(req)
	log.Print(vars)
	id := bson.ObjectIdHex(vars["id"])
	log.Print("ID",id)
	var dataResource CompanyResource
	// Decode the incoming Company json

	err := json.NewDecoder(req.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid company data",
			500,
		)
		return
	}
	company := &dataResource.Data
	company.CompanyID = id
	context := NewContext()
	log.Print("Created context update...")

	defer context.Close()
	c := context.DbCollection("companies")

	// Update an existing Company document
	repo := &data.CompanyRepository{c}

	// Updating an existing company document
	if err := repo.UpdateCompany(company);err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}else {
		w.WriteHeader(http.StatusNoContent)
	}
}


func GetCompanies(w http.ResponseWriter, req *http.Request){
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("companies")
	CompaniesRep := &data.CompaniesRepository{c}
	companies := CompaniesRep.GetAllCompanies()
	j, err := json.Marshal(CompaniesResource{Data: companies})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An Unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetCompanyByID(w http.ResponseWriter, req *http.Request){
	// get the id from the incoming url
	log.Print("From => GetCompanyByID()")
	vars := mux.Vars(req)
	id := vars["id"]
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("companies")
	repo := &data.CompanyRepository{c}
	company, err := repo.GetCompanyById(id)
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
	if j, err := json.Marshal(company); err != nil {
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
func DeleteCompany(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	log.Print(vars)
	id := vars["id"]
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("companies")
	repo := &data.CompanyRepository{c}

	// Delete an existing company document
	err := repo.DeleteCompany(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred, while deleting the company",
			500,
		)
		return

	}
	w.WriteHeader(http.StatusNoContent)
}

func SendInvestorInvitation(w http.ResponseWriter, req *http.Request) {

	iss := req.Context().Value("iss").(string)
	email := req.Context().Value("email").(string)
	id := req.Context().Value("id").(string)
	log.Printf("iss claims: %#v",iss)

	type Invitation struct {
		Email string `json:"email"`
		Note  string `json:"note"`
		CompanyID string `json:"companyid"`
	}

	Inv := Invitation{}
	if req.Body == nil {
		http.Error(w, "Please send a request data.", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&Inv)
	if err != nil {
		log.Println(err)
	}
	log.Println(Inv)

	auth := smtp.PlainAuth(
		"",
		"concludeas@gmail.com",
		"techrepublic.com",
		"smtp.gmail.com",
	)

	emailBody := fmt.Sprintf(`
		Body: %s <br>
		User Email: %s <br>
		User ID: %s
		Company Invite Link: http://127.0.0.1:3000/auth?companyid=%s
	`, Inv.Note, email, id, Inv.CompanyID)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"concludeas@gmail.com",
		[]string{Inv.Email},
		[]byte(emailBody),
	)
	if err != nil {

		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := make(map[string]interface{})
	message["message"] = "Your invitation has been sent!"
	message["code"] = http.StatusOK
	json.NewEncoder(w).Encode(message)
	//w.Write([]byte("You invitation has been sent!"))
}
/*
func AddInvestorToCompany(w http.ResponseWriter, req *http.Request){
	// User ID
	// Company ID
	// InvestmentID
	collection.Update(
		bson.M{
			"_id":bson.ObjectIdHex(id),
		},
		bson.M{
			"$push":
			bson.M{"company.$.investors":
			bson.M{
				"investorid": x,
				"y":y,
			}
			}
		}
	)

	var dataResource CompanyResource

	// Decoding the incoming company
	err := json.NewDecoder(req.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid company data",
			500,
		)
		return
	}

	company := &dataResource.Data
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("companies")
	repo := &data.CompanyRepository{c}

	// insert a company document
	repo.CreateCompany(company)
	if j, err := json.Marshal(CompanyResource{Data: *company}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return

	} else {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		w.Write(j)
	}

	)


}*/
func GetCompanyByUser(w http.ResponseWriter, r *http.Request){
	// Get id from the incoming url

	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()

	c:= context.DbCollection("companies")
	repo := &data.CompanyRepository{c}
	companies := repo.GetCompanyByUser(user)
	j, err := json.Marshal(CompaniesResource{Data: companies})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}


// Add Shareholder to company
func AddShareholder(w http.ResponseWriter, req *http.Request){
	var dataResource CompanyResource

	// Decoding the incoming company
	err := json.NewDecoder(req.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid company data",
			500,
		)
		return
	}

	company := &dataResource.Data
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("companies")
	repo := &data.CompanyRepository{c}

	// insert a company document
	repo.CreateCompany(company)
	if j, err := json.Marshal(CompanyResource{Data: *company}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return

	} else {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		w.Write(j)
	}

}
