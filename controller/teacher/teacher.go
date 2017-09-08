package teacher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mvcapp/controller"
	"mvcapp/model/teacher"
	"net/http"
)

func Load(router *mux.Router) {
	router.HandleFunc("/teacher/add", Add).Methods("POST").Name("TeacherAdd")
	router.HandleFunc("/teacher/delete/{id}", Delete).Methods("DELETE").Name("TeacherDelete")
	router.HandleFunc("/teacher/update/{id}", Update).Methods("PUT").Name("TeacherUpdate")
	router.HandleFunc("/teacher/get/{id}", View).Methods("GET").Name("TeacherView")
}

/*
Add method adds a new record in teachers collection

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
		errMsg := controller.Error_Status{Message: err.Error()}
		raw_err, _ := json.Marshal(errMsg)
		res.Write(raw_err)
		//http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	obj := teacher.Teacher{}
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
Add method updated a record in teachers collection based upon id

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
	obj := teacher.Teacher{}

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
Delete method deletes a record from the teachers collection
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
	obj := teacher.Teacher{ID: id}
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

}

/*

View method gets a record from teachers collection based upon id

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
	obj := teacher.Teacher{}
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

/*
func Error(res *http.ResponseWriter,err error){
	fmt.Printf("%s", err)
	errMsg := controller.Error_Status{Message:err.Error()}
	raw_err,_ := json.Marshal(errMsg)

	http.Error(res, err.Error(), http.StatusInternalServerError)
}*/
