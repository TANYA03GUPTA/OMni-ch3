package main

import (
	"html/template"
	"net/http"
	"time"
)

type pgData struct{
	title string 
	num []int
}

func fetchdt(ch chan []int){
	time.Sleep(2 * time.Second)
	ch <- []int{1,2,3,4,5}
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	tmp2, err := template.ParseFiles("t1/html")
	tmpl, err := template.New("home").Parse(
		`<html>
		    <head><title>{{ .Title }}</title></head>
            <body>
                <h1>{{ .Title }}</h1>
                <ul>
                    {{ range .Numbers }}
                        <li>{{ . }}</li>
                    {{ end }}
                </ul>
            </body>
        </html>
		`)
	if err != nil{
		http.Error(w, "error parsing template", http.StatusInternalServerError)
		return 
	}
	ch := make(chan []int)
	go fetchdt(ch)

	numb := <- ch
	data := pgData{
		title: "Home Page",
		num: numb,
	}
	tmpl.Execute(w, data)
}

func main(){
	http.HandleFunc("/",homeHandler)
	http.ListenAndServe(":5692",nil)
}