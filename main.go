package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"os"
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
// Get the Port from the environment so we can run on Heroku
func GetPort() string {
 	var port = os.Getenv("PORT")
 	// Set a default port if there is nothing in the environment
 	if port == "" {
 		port = "4747"
 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", IntroductionHandler)
	http.HandleFunc("/Methodology",MethodologyHandler)
	//http.HandleFunc("/poster",PosterHandler)


	fs := http.FileServer(http.Dir(templatesdirectoryPath))
	http.Handle(imagesPath, fs)

	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

