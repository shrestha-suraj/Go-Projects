package main

import (
	"log"
	"net/http"
	"text/template"
)

type Course struct {
	Code     string
	Title    string
	Semester string
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func GoTemp(w http.ResponseWriter, r *http.Request) {
	C := []Course{}
	C = append(C, (Course{Code: "CSCI 111", Title: "Computer Science I", Semester: "Fall & Spring"}))
	C = append(C, (Course{Code: "CSCI 112", Title: "Computer Science II", Semester: "Fall & Spring"}))
	C = append(C, (Course{Code: "CSCI 211", Title: "Computer Science III", Semester: "Fall & Spring"}))
	C = append(C, (Course{Code: "CSCI 223", Title: "Computer Organization & Assembly Language", Semester: "Fall & Spring"}))
	C = append(C, (Course{Code: "CSCI 300", Title: "Social Responsibility in Computer Science", Semester: "Fall"}))
	C = append(C, (Course{Code: "CSCI 311", Title: "Models of Computation", Semester: "Fall"}))
	C = append(C, (Course{Code: "CSCI 387", Title: "Software Design and Development", Semester: "Spring"}))
	C = append(C, (Course{Code: "CSCI 423", Title: "Introduction to Operating System", Semester: "Fall"}))
	C = append(C, (Course{Code: "CSCI 433", Title: "Algorithm and Data Structure Analysis", Semester: "Spring"}))
	C = append(C, (Course{Code: "CSCI 450", Title: "Organization of Programming Languages", Semester: "Fall"}))
	C = append(C, (Course{Code: "CSCI 487", Title: "Senior Project", Semester: "Fall & Spring"}))
	tmpl.ExecuteTemplate(w, "GoTemplatesHW", C)
}
func main() {
	log.Println("Server started on: http://localhost:8123")
	http.HandleFunc("/", GoTemplatesHW)
	http.ListenAndServe(":8123", nil)
}
