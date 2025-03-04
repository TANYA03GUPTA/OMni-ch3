package main

import (
	"database/sql"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

const staticPath string = "static/"
type WebPage struct {
	Title string
	Contents string
	Connection *sql.DB
	}

	type customRouter struct {
	}
	func serveDynamic() {
	}
	func serveRendered() {
	}
	func serveStatic() {
	}	

func (customRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path;
		staticPatternString := "static/(.*)"
		templatePatternString := "template/(.*)"
		dynamicPatternString := "dynamic/(.*)"
		staticPattern := regexp.MustCompile(staticPatternString)
		templatePattern := regexp.MustCompile(templatePatternString)
		dynamicDBPattern := regexp.MustCompile(dynamicPatternString)
		if staticPattern.MatchString(path) {
		serveStatic()
		page := staticPath + staticPattern.ReplaceAllString(path,
		"${1}") + ".html"
		http.ServeFile(rw, r, page)
		}else if templatePattern.MatchString(path) {
		serveRendered()
		urlVar := templatePattern.ReplaceAllString(path, "${1}")
		page.Title = "This is our URL: " + urlVar
		customTemplate.Execute(rw,page)
		}else if dynamicDBPattern.MatchString(path) {
		serveDynamic()
		page = getArticle(1)
		customTemplate.Execute(rw,page)
	}
}	
