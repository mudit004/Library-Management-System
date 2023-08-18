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

	r.HandleFunc("/", controller.Login).Methods("GET")
	r.HandleFunc("/admin", controller.Admin).Methods("GET")
	r.HandleFunc("/register", controller.Register).Methods("GET")
	r.HandleFunc("/addBook", controller.AddBookPage).Methods("GET")
	r.HandleFunc("/returnBook", controller.ReturnBookLoader).Methods("GET")
	r.HandleFunc("/requestBook", controller.RequestBookLoader).Methods("GET")
	r.HandleFunc("/pendingRequest", controller.PendingRequestLoader).Methods("GET")
	r.HandleFunc("/internalServerError", controller.ServerUnavailable).Methods("GET")
	r.HandleFunc("/requestManagement", controller.RequestManagementPage).Methods("GET")

	r.HandleFunc("/logout", controller.Logout).Methods("POST")
	r.HandleFunc("/login", controller.LoginRequest).Methods("POST")
	r.HandleFunc("/adminLogin", controller.AdminLogin).Methods("POST")
	r.HandleFunc("/admin/add", controller.AddBookHandler).Methods("POST")
	r.HandleFunc("/registerUser", controller.RegisterUser).Methods("POST")
	r.HandleFunc("/admin/remove", controller.RemoveBookHandler).Methods("POST")
	r.HandleFunc("/incrementBook", controller.IncrementBookHandler).Methods("POST")
	r.HandleFunc("/decrementBook", controller.DecrementBookHandler).Methods("POST")
	r.HandleFunc("/userSide/checkin", controller.RequestBookHandler).Methods("POST")
	r.HandleFunc("/userSide/checkout", controller.ReturnBookHandler).Methods("POST")
	r.HandleFunc("/admin/rejectRequest", controller.RejectRequestHandler).Methods("POST")
	r.HandleFunc("/admin/acceptRequest", controller.AcceptRequestHandler).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(controller.NotFoundHandler)

	http.ListenAndServe(":8000", r)
}
