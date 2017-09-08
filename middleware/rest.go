package middleware

import (
	"github.com/gorilla/mux"
	"mvcapp/controller/student"
	"mvcapp/controller/teacher"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	student.Load(router)
	teacher.Load(router)
	return router
}
