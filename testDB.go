package main

//import packages
import(
	"database/sql"
	"database1"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)


func dbConn() (db *sql.DB) {
  //Using your variables in database.go, establish a connection with your database
	user:=database.Usr()
	password:=database.Pwd()
	db,err:=sql.Open("mysql",user+":"+password+"@/"+user)
  //Call the handler db  - this is what you return below

  if err != nil {
        panic(err.Error())
  }
  return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func TestDBConnect(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
	
	//NOTE:  This select statement assumes you still have a Formula table
    _, err := db.Query("SELECT * FROM Formula")
    if err != nil {
        panic(err.Error())
    }
	tmpl.ExecuteTemplate(w, "testDBConnect.html", nil)
	defer db.Close()
}

func main() {
	//Replace 1111 with your localhost
	//Replace **********  with your webid
    log.Println("Server started on: http://localhost:8123")
    http.HandleFunc("/", TestDBConnect)
    http.Handle("/home/sshrest4/static/", http.StripPrefix("/home/sshrest4/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8123", nil)
}

