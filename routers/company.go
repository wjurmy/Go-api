package routers

import (
	"github.com/wjurmy/havaard/common"
	"github.com/wjurmy/havaard/controllers"

	// Third Party libraries
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

)

/*			Routes for the Company Resource in Company.go
		----------------------------------------
SetUserRoutes function receives a pointer to the Gorilla mux
router object ( mux.Router ) as an argument pointer of the
mux.Router object. Two routes are specified: for registering
a new user login to the system. Application handler functions
called from the controllers package.
*/
func SetCompanyRoutes(router *mux.Router) *mux.Router{
	companyRouter := mux.NewRouter()
	companyRouter.HandleFunc("/company",controllers.RegisterCompany).Methods("POST")
	companyRouter.HandleFunc("/company/{id}", controllers.UpdateCompany).Methods("PUT")
	companyRouter.HandleFunc("/company",controllers.GetCompanies).Methods("GET")
	companyRouter.HandleFunc("/company/users/{id}", controllers.GetCompanyByUser).Methods("GET")
	companyRouter.HandleFunc("/company/{id}",controllers.GetCompanyByID).Methods("GET")
	companyRouter.HandleFunc("/company/{id}",controllers.DeleteCompany).Methods("DELETE")
	companyRouter.HandleFunc("/company/invitation",controllers.SendInvestorInvitation).Methods("POST")
	companyRouter.HandleFunc("/company/shareholder",controllers.AddShareholder).Methods("POST")

	//companyRouter.HandleFunc("/company/shareholder",controllers.AddShareholder).Methods("POST")

	/*
	Adding route-specific middelware
	--------------------------------
	Authorization middelware (negroni ) for the route path "/company"
	restrics access only to authenticated users. In the
	SetCompanyRoutes funtion. A new router instance of mux router
	is created for the "/company" resources specified, and the
	autorization middleware is wrapped into the handler functions of the routes path "/path"

	*/
	router.PathPrefix("/company").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(companyRouter),
	))
	return router
}
