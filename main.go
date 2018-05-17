package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)
const (
	templatesdirectoryPath="templates"
	cssPath="/css/"
	jsPath="/js/"
	imagesPath="/images/"

)
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "index", data)
}


func IntroductionHandler(w http.ResponseWriter, r *http.Request) {

	generateHTML(w,nil,"index","Introduction")
}


func MethodologyHandler(w http.ResponseWriter, r *http.Request){

	generateHTML(w,nil,"index","methodology")

}

//func PosterHandler(w http.ResponseWriter, r *http.Request){

//	generateHTML(w,nil,"index","poster")

//}

func main() {
	http.HandleFunc("/", IntroductionHandler)
	http.HandleFunc("/Methodology",MethodologyHandler)
	//http.HandleFunc("/poster",PosterHandler)


	fs := http.FileServer(http.Dir(templatesdirectoryPath))
	http.Handle(imagesPath, fs)

	log.Fatal(http.ListenAndServe("dfswebapp.herokuapp.com:8080", nil))
}

