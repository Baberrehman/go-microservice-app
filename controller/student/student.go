package student

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mvcapp/model/student"
	"net/http"
)

func Load(router *mux.Router) {
	router.HandleFunc("/student/add", Add).Methods("POST").Name("Add")
	router.HandleFunc("/student/delete/{id}", Delete).Methods("DELETE").Name("Delete")
	router.HandleFunc("/student/update/{id}", Update).Methods("PUT").Name("Update")
	router.HandleFunc("/student/get/{id}", View).Methods("GET").Name("View")
}

/*
Add method adds a new record in students collection

Sample input:
{
	"id":"1",
	"name":"jonh",
	"password":"123"
}
*/

func Add(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	obj := student.Student{}
	err = json.Unmarshal(contents, &obj)
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = obj.Insert()
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(`{"msg" :"data inserted successfully"}`))
}

/*
Update method updates a student in students collection based upon id as primary key

Sample input:
{
	"id":"1",
	"name":"jonh",
	"password":"123"
}
*/

func Update(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Printf("%s", "id not found in query param")
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, errors.New("ID not found in query param").Error(), http.StatusInternalServerError)
		return
	}
	obj := student.Student{}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
		/*errMsg := controller.Error_Status{Message:err.Error()}
		raw_err,_ := json.Marshal(errMsg)
		res.Write(raw_err)*/
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(contents, &obj)
	_, err = obj.Update(id)
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(`{"msg" :"data updated successfully"}`))
}

/*
Delete method deletes record from the students collection based upon id

Sample input:
{
	"id":"1"
}
*/

func Delete(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Printf("%s", "id not found in query param")
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, errors.New("ID not found in query param").Error(), http.StatusInternalServerError)
		return
	}
	obj := student.Student{ID: id}
	_, err := obj.Delete()
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(`{"msg" :"data deleted successfully"}`))

	//c := flight.Context(w, r)

	//c.View.New("about/index").Render(w, r)
}

/*
Delete method deletes a record from the students collection based upon id

Sample input:
{
	"id":"1",
	"name":"jonh",
	"password":"123"
}
*/

func View(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Printf("%s", "id not found in query param")
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, errors.New("ID not found in query param").Error(), http.StatusInternalServerError)
		return
	}
	obj := student.Student{}
	_, err := obj.Get(id)
	if err != nil {
		fmt.Printf("%s", err)
		//errMsg := controller.Error_Status{Message:err.Error()}
		//raw_err,_ := json.Marshal(errMsg)
		//res.Write(raw_err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	raw_student, _ := json.Marshal(obj)
	fmt.Fprint(res, string(raw_student))
}
