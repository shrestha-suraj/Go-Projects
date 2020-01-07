package main

import(
	"log"
	"net/http"
	"text/template"
)

type SportsLeagues struct{
	Sport string
	Championship string
}
var tmpl = template.Must(template.ParseGlob("templates/*"))

func GoTemp(w http.ResponseWriter,r *http.Request){
	SL:=[]SportsLeagues{}
	SL=append(SL, (SportsLeagues{Sport: "Football",Championship: "Superbowl"}))
	SL=append(SL, (SportsLeagues{Sport: "NBA",Championship: "Championship Series"}))
	SL=append(SL, (SportsLeagues{Sport: "MLB",Championship: "World Series"}))
	SL=append(SL, (SportsLeagues{Sport: "FIFA",Championship: "World Cup"}))
	tmpl.ExecuteTemplate(w,"GoTemp",SL)
}
func main(){
	log.Println("Server started on: http://localhost:8123")
	http.HandleFunc("/",GoTemp)
	http.ListenAndServe(":8123",nil)
}