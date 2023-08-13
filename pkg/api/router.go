package api

import (
	"lms/pkg/controller"
	"lms/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.Use(utils.ValidateJWT)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	r.HandleFunc("/browse", controller.Browse).Methods("GET")
	r.HandleFunc("/dashboard", controller.Dashboard).Methods("GET")

	r.HandleFunc("/", controller.Login).Methods("GET")
	r.HandleFunc("/login", controller.LoginRequest).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("GET")
	r.HandleFunc("/registerUser", controller.RegisterUser).Methods("POST")
	r.HandleFunc("/admin/add", controller.AddBookHandler).Methods("POST")
	r.HandleFunc("/admin/rem", controller.RemBookHandler).Methods("POST")
	r.HandleFunc("/admin/acceptReq", controller.AcceptRequestHandler).Methods("POST")
	r.HandleFunc("/admin/rejectReq", controller.RejectRequestHandler).Methods("POST")
	r.HandleFunc("/user_side/checkin", controller.RequestBookHandler).Methods("POST")
	r.HandleFunc("/user_side/checkout", controller.ReturnBookHandler).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")

	http.ListenAndServe(":8000", r)
}
