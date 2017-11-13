package models

import (
	"time"
	//_ "github.com/jmoiron/sqlx"
	//_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
)

type (

	User struct {
		Id			bson.ObjectId	`bson:"_id" json:"id" db:"_id"`
		FirstName	 string 		`json:"FirstName,omitempty"			db:"FirstName"`
		LastName	 string			`json:"LastName,omitempty" 			db:"LastName"`
		Email		 string			`json:"Email,omitempty" 			db:"Email"`
		HashPassword []byte			`json:"-,omitempty" 				db:"Password"`
		Password	 string			`json:"Password,omitempty" 			db:"-"`
		CreatedOn	 time.Time		`json."CreatedOn,omitempty" 		db:"CreatedOn"`
		UpdatedOn	 time.Time		`json."UpdatedOn,omitempty" 		db:"UpdatedOn"`
		Companies    []string `json:"-"`
		Role struct {
			Admin	 	 bool 		`json:"Admin,omitempty"				db:"Admin"`
			Editor	 	 bool 		`json:"Editor,omitempty"		    db:"Editor"`
			Investor	 bool 		`json:"Investor,omitempty"			db:"Investor"`
		}

		Address struct {
			Id			bson.ObjectId	`bson:"_AddressID,omitempty" json:"id"	db:"AddressID"`
			City 		string 			`json:"city,omitempty"`
			Street 		string 			`json:"state,omitempty"`
			PostBox 	string			`json:"state,omitempty"`
		}




	}

	Company struct {
		Id		 		bson.ObjectId 	`bson:"_id,omitempty" 				db:"_id"`
		CompanyID 		bson.ObjectId 	`bson:"_CompanyID,omitempty" 		db:"_CompanyID"`
		CompanyName 	string			`json:"CompanyName,omitempty" 		db:"CompanyName"`
		CompanyOrgNr	string 			`json:"CompanyOrgNr,omitempty" 		db:"CompanyOrgNr"`
		CompanyAddress	string 			`json:"CompanyAddress,omitempty" 	db:"CompanyAddress"`
		CreatedOn		time.Time 		`json:"CreatedOn,omitempty" 		db:"CreatedOn"`
		Postbox			string 			`json:"Postbox,omitempty" 			db:"Postbox"`
		UpdatedOn	 	time.Time		`json."UpdatedOn,omitempty" 		db:"UpdatedOn"`
		CreatedBy 		string 			`json:"CreatedBy,omitempty"			db:"CreatedBy"`
		Description 	string 			`json:"Description,omitempty"		db:"Description"`
		Email		 	string			`json:"Email,omitempty" 			db:"Email"`
		Telefon		 	string			`json:"Telefon,omitempty" 			db:"Telefon"`


		Investor []struct {
			UserID 			string 			`json:"UserID,omitempty"				db:"UserID"`
			InvestmentID 	string 			`json:"InvestmentID,omitempty"			db:"InvestmentID"`
		}

		Investment []struct {
			InvestorType 	string 			`json:"InvestorType,omitempty"			db:"InvestorType"`
			InvestorID 		string			`json:"InvestorID,omitempty"			db:"InvestorID"`
			CompanyID 		bson.ObjectId 	`bson:"_CompanyID,omitempty" 			db:"_CompanyID"`
			RaisedAmount 	int 			`json:"RaisedAmount,omitempty"			db:"RaisedAmount"`
		}
	}






/*	Response struct {
		Data string `json:"data"`
	}

	Token struct {
	Token string `json:"token"`
	}
*/
)