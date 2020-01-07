package main

import (
    "database/sql"
	"fmt"
    "log"
    "net/http"
    "text/template"
	"database/database"

    _ "github.com/go-sql-driver/mysql"
)

type Person struct {
    PersonID  int
    FirstName    string
    LastName  string
    Birthdate  string
    BirthCity  string
    BirthState string
    Region string
}

func dbConn() (db *sql.DB) {
	//Copy over the body of your dbConn() from the PresidentsDB lab/homework
	user:= database.Usr();
	password:=database.Pwd();
	db, err := sql.Open("mysql", user+":"+password+"@/"+user)
	if err != nil {
        panic(err.Error())
	}
	//UNCOMMENT LINE BELOW ONCE FUNCTION IS IMPLEMENTED
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    //Connect to your database using "db" as the name your handler
	db:=dbConn()
	//Select from people table - you only need PersonID, FirstName, and LastName
	row, err := db.Query("SELECT PersonID, FirstName, LastName FROM people")
	if err !=nil{
		panic(err.Error())
	}
	//Create a "results" slice of Person, adding each row from your query to this slice
	results:= [] Person{}
	for row.Next(){
		var idnumber int
		var firstname string
		var lastname string
		err = row.Scan(&idnumber,&firstname,&lastname)
		if err!=nil{
			panic(err.Error())
		}
		results=append(results, Person{idnumber,firstname,lastname,"","","",""})
	}
	
	
	//Execute template, sending "results" to index.html and defer closing db
	//UNCOMMENT LINES BELOW ONCE FUNCTION IS IMPLEMENTED
    tmpl.ExecuteTemplate(w, "index.html", results)
    defer db.Close()
}

func Create(w http.ResponseWriter, r *http.Request) {
    //Execute template that will provide a form for adding a person
	//UNCOMMENT LINE BELOW WHEN INSTRUCTED IN LAB
	tmpl.ExecuteTemplate(w, "create.html", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
    //This function is called AFTER the form in create.html is submitted
	//Verify there is a POST 
	db:=dbConn()
	if r.Method=="POST"{
		firstName:=r.FormValue("FirstName")
		lastName:=r.FormValue("LastName")
		birthDate:=r.FormValue("BirthDate")
		birthCity:=r.FormValue("BirthCity")
		birthState:=r.FormValue("BirthState")
		region:=r.FormValue("Region")
		stmt, err := db.Prepare("INSERT INTO people (FirstName,LastName,Birthdate,BirthCity,BirthState,Region) VALUES(?,?,?,?,?,?)")
		if err!=nil{
			panic(err.Error())
		}
	stmt.Exec(firstName,lastName,birthDate,birthCity,birthState,region);
	log.Println("ADDED: Name: " + firstName + " " + lastName)
	}
	
	//This function will add a person's information so use the "POST" fields from the form
	//Look back at your lab and homework on PresidentsDB for syntax to process form fields
	
	//Prepare a statement to Insert Into people
	//Execute statement using form values	
		//Log updated information - will print to Command Line
		//Change my variable names firstName and lastName to what you defined then UNCOMMENT statement below	
		

		
		
	//Defer closing database
	//UNCOMMENT LINES BELOW ONCE FUNCTION IS IMPLEMENTED	
	defer db.Close()

	//****************Redirect to index.html "/"
	http.Redirect(w, r,"/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	//This function is called AFTER a link to a specific person has been selected on the index.html page
	
    //Connect to your database using "db" as the name of your handler
	db:=dbConn()
	//This function is editing a person so "Get" their "id" from the query string
	personId:=r.URL.Query().Get("id")
	row, err :=db.Query("SELECT * FROM people WHERE PersonID=?",personId)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println("Pass1")
	fmt.Println("Pass2")
	onePerson:= [] Person{}
	for row.Next(){
		var idnumber int
		var firstname string
		var lastname string
		var birthdate string
		var birthcity string
		var birthstate string
		var region string
		err = row.Scan(&idnumber,&firstname,&lastname,&birthdate,&birthcity,&birthstate,&region)
		if err!=nil{
			panic(err.Error())
		}
		onePerson=append(onePerson, Person{idnumber,firstname,lastname,birthdate,birthcity,birthstate,region})
	}
	
	
	//Using this "id", query your database to get this person's information
	//Create a "onePerson" slice of Person, adding the one row from your query to this slice
	

	
   //Execute template, sending results to edit.html and defer closing db
   //UNCOMMENT LINES BELOW ONCE FUNCTION IS IMPLEMENTED
   tmpl.ExecuteTemplate(w, "edit.html", onePerson)
   defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
    //This function is called AFTER the form in edit.html is submitted
	
	//Connect to your database using "db" as the name of your handler
	db:=dbConn()
	if r.Method=="POST"{
		pID:=r.FormValue("PersonID")
		firstName:=r.FormValue("FirstName")
		lastName:=r.FormValue("LastName")
		birthDate:=r.FormValue("BirthDate")
		birthCity:=r.FormValue("BirthCity")
		birthState:=r.FormValue("BirthState")
		region:=r.FormValue("Region")
		stmt, err := db.Prepare("UPDATE people SET FirstName=?, LastName=?, BirthDate=?,BirthCity=?, BirthState=?, Region=? WHERE PersonID=?")
		if err!=nil{
			panic(err.Error())
		}
	stmt.Exec(firstName,lastName,birthDate,birthCity,birthState,region,pID);
	log.Println("UPDATE: Name: " + firstName + " " + lastName + " | ID: " + pID)
	}
	//This function will update a person's information so "Post" fields from the form
	//Look back at your lab and homework on PresidentsDB for syntax to process form fields
	
	//Prepare a statement to Update people
	//Execute statement using form values



	
		//Log updated information - will print to Command Line
		//Change my variable names firstName and lastName to what you defined then UNCOMMENT statement below
		
	
 
	//Defer closing database 

	//UNCOMMENT LINES BELOW ONCE FUNCTION IS IMPLEMENTED
	defer db.Close()
	
	//****************Redirect to index.html "/"
	http.Redirect(w, r,"/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//This function is called AFTER a link to a specific person has been selected on the index.html page
	
    //Connect to your database using "db" as the name of your handler
	db:=dbConn()
	//This function is deleting a person so "Get" their "id" from the query string
	pID:=r.URL.Query().Get("id")
	//Using this "id", prepare a statement to Delete From people
	//Execute statement using "id" as the criterion
	db.Query("DELETE FROM people WHERE PersonID=?",pID)
	log.Println("DELETE")
	
	//Log delete
    

	//Defer closing database 

	//UNCOMMENT LINES BELOW ONCE FUNCTION IS IMPLEMENTED
	defer db.Close()

	//****************Redirect to index.html "/"
	http.Redirect(w,r,"/",301)
}

func main() {
    log.Println("Server started on: http://localhost:8123")    //*****Replace 1111 with your port
    http.HandleFunc("/", Index)
    http.HandleFunc("/create.html", Create)
    http.HandleFunc("/add", Add)
    http.HandleFunc("/edit.html", Edit)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/index.html", Delete)
     //Replace with your Handle - http.Handle
	 http.Handle("/home/sshrest4/static/", http.StripPrefix("/home/sshrest4/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8123", nil)			//**********Replace 1111 with your port
}
