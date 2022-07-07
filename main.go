package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"problem1/db"
	"problem1/entities"

	"github.com/go-chi/chi"
)

var company entities.Company

func main() {
	db.InitDatabase()
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Post("/postLocation", PostLocation)
		r.Get("/getLocations", GetLocations)
		r.Post("/postEmployee", PostEmployees)
		r.Get("/getEmployees", GetEmployees)
		r.Post("/postCompany", PostCompany)
		r.Get("/getCompanies", GetCompanies)

	})
	http.ListenAndServe(":8080", router)
}

func PostLocation(w http.ResponseWriter, r *http.Request) {
	location := entities.Location{}
	body := r.Body
	desObj, _ := io.ReadAll(body)
	json.Unmarshal(desObj, &location)
	db.GetDB().Create(&location)
	fmt.Fprint(w, location)
}
func PostEmployees(w http.ResponseWriter, r *http.Request) {
	employee := entities.Employee{}
	body := r.Body
	desObj, _ := io.ReadAll(body)
	json.Unmarshal(desObj, &employee)
	db.GetDB().Create(&employee)
	fmt.Fprint(w, employee)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []entities.Employee{}
	db.GetDB().Find(&employees)
	fmt.Fprint(w, employees)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	locations := []entities.Location{}
	db.GetDB().Find(&locations)
	fmt.Fprint(w, locations)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	companies := []entities.Company{}
	db.GetDB().Preload("Location").Find(&companies)
	arr:= []string{}
	for _, company := range companies {
		arr = append(arr,company.Name)
		arr = append(arr, company.Location.City)
		
	}
	
	fmt.Fprint(w, arr)

}
func PostCompany(w http.ResponseWriter, r *http.Request) {
	company := entities.Company{}
	body := r.Body
	desObj, _ := io.ReadAll(body)
	json.Unmarshal(desObj, &company)
	db.GetDB().Create(&company)
	fmt.Fprint(w, company)
}
