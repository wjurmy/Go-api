package data

import (
	"github.com/wjurmy/havaard/models"
	//"github.com/jmoiron/sqlx"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"
	"log"

)

type CompanyRepository struct {
	C *mgo.Collection
}

type CompaniesRepository struct {
	CompaniesRep *mgo.Collection
}

func (r *CompanyRepository) CreateCompany(company *models.Company) error {
	log.Print(company.CompanyName)
	obj_id := bson.NewObjectId()
	company.CompanyID = obj_id
	company.CreatedOn = time.Now()
	err := r.C.Insert(&company)
	return err
}
// This method returns all the companies
func (r *CompaniesRepository) GetAllCompanies() []models.Company{
	var companies []models.Company
	iter := r.CompaniesRep.Find(nil).Iter()
	result := models.Company{}
	for iter.Next(&result){
		companies = append(companies, result)

	}
	return companies

}

func (r *CompanyRepository) GetCompanyById(id string) (company models.Company, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&company)
	return
}

//TODO can't delete the company
func (r *CompanyRepository) DeleteCompany(id string) error{
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *CompanyRepository ) UpdateCompany(company *models.Company) error {
	// partial update on MongoDB

	err := r.C.Update(bson.M{"_id": company.CompanyID},
		bson.M{"$set": bson.M{
			"companyname": company.CompanyName,
			"companyorgnr": company.CompanyOrgNr,
			"companyaddress": company.CompanyAddress,
			"updatedon": time.Now(),
		}})
	return err

}

func (r *CompanyRepository) GetCompanyByUser(user string) []models.Company {
	var companies []models.Company
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Company{}
	for iter.Next(&result){
		companies = append(companies, result)
	}
	log.Print(companies)
	return companies
}
