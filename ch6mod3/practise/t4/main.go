package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func loginhandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"welcome to the login apde !")
}

func homelander(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"welcome to the home page !")
}

func UserHandler(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Fprintf(w,"helloe user %s ", username)
	user1 := strings.TrimPrefix(r.URL.Path, "/user/")
	fmt.Fprintf(w,"user name is %s ", user1)
}

func searchhandler(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	keyword := q.Get("keyword")
	fmt.Fprintf(w,"search keyword is %s ", keyword)
}

func shophandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	cat := vars["category"]
	id := vars["id"]
	fmt.Fprintf(w, "Category is : %s, Product : %s", cat,id)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/home",homelander)
	router.HandleFunc("/login",loginhandler)
	router.HandleFunc("/user/{username}", UserHandler)
	router.HandleFunc("/search",searchhandler)
	router.HandleFunc("/shop/{category}/product/{id}", shophandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	fmt.Println("server is up & running")
	http.ListenAndServe(":5860",router)
}