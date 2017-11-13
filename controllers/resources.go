package controllers

import (
	"github.com/wjurmy/havaard/models"
)

type (
	CompanyResource struct {
		Data models.Company `json:"data"`
	}

	// For Get - /tasks
	CompaniesResource struct {
		Data []models.Company `json:"data"`
	}

	UserResource struct {
		CompanyID string `json:"companyid"`
		Data models.User `json:"data"`
	}

	/*AddressResource struct {
		Data models.User.Address `json:"data"`
	}*/
	
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	
	
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	
	LoginModel struct {
		Email		string `json:email`
		Password	string `json:"password"`
	}
	
	AuthUserModel struct {
		User	models.User `json:"user"`
		Token 	string 		`json:"token"`
	}

)