package main

import(
	"log"
	"net/http"
	"text/template"
	"strconv"
	"database/sql"
	"database/database"
	_ "github.com/go-sql-driver/mysql"

)

func dbConn() (db *sql.DB) {
  //Using your variables in database.go, establish a connection with your database
  user:= database.Usr();
  password:=database.Pwd();
  //Call the handler db  - this is what you return below
  db, err := sql.Open("mysql", user+":"+password+"@/"+user)
  if err != nil {
        panic(err.Error())
  }
  return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func ChoosePresident2019(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

	//NOTE:  This select statement assumes you still have a Formula table
    row, err := db.Query("SELECT * FROM presidents")
    if err != nil {
        panic(err.Error())
    }

	type President struct {
	 Num int64
	 LName string
	 FName string
	 MInitial string
	 State string
	 Party string
	 Start int64
	 End int64
	 Term string
    }
	USPresidents:= []President{}
	for row.Next(){
		//Create string and int64 variables to be used as references to
		//each of the attributes – 9 in all. Refer to the sql script.
		var number int64
		var lname string
		var fname string
		var mInitial string
		var state string
		var party string
		var start int64
		var end int64
		var term string

		//Scan each row assigning to the variables defined above. Note
		//that below I use var1, var2, … – you should use more meaningful
		//names – specifically the names you defined for each attribute
		//Recall that & in Go is a reference. We use & because Scan assigns
		//the result of this row’s data to each of the variable references
		err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
		if err != nil {
			panic(err.Error())
		}
		//Append to your USPresidents slice – use the variables defined
		//above to create a President element added to the USPresidents
		//struct. Note that you DO NOT use & here because we want
		//the value assigned to the variable, not its reference
		USPresidents = append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
	}

	tmpl.ExecuteTemplate(w, "choosePresidentDB2019.html", USPresidents)
	defer db.Close()
}
func ListPresident2019(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
	//NOTE:  This select statement assumes you still have a Formula table
	//add codes here
	//define USPResidents slice
	type President struct {
	 Num int64
	 LName string
	 FName string
	 MInitial string
	 State string
	 Party string
	 Start int64
	 End int64
	 Term string
    }
	USPresidents:= []President{}

	if r.Method=="POST"{
		IDStr:= r.FormValue("ID")
		if IDStr!=""{
			ID,_ :=strconv.ParseInt(IDStr, 0, 64)
			row, err := db.Query("SELECT * FROM presidents WHERE number = ?", ID)
			if err != nil {
				panic(err.Error())
			}
			for row.Next(){
				var number int64
				var lname string
				var fname string
				var mInitial string
				var state string
				var party string
				var start int64
				var end int64
				var term string
				err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
				if err != nil {
					panic(err.Error())
				}
				USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
			}
		}else {
			first_name := r.FormValue("FName")
			if first_name != "" {
				fName :="%"+first_name+"%"
				row, err := db.Query("SELECT * FROM presidents WHERE fname LIKE ?",fName)
				if err != nil {
					panic(err.Error())
				}
				for row.Next(){
					var number int64
					var lname string
					var fname string
					var mInitial string
					var state string
					var party string
					var start int64
					var end int64
					var term string
					err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
					if err != nil {
						panic(err.Error())
					}
					USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
			}
				
			}else {
				last_name := r.FormValue("LName")
				if last_name != "" {
					lName :="%"+last_name+"%"
					row, err := db.Query("SELECT * FROM presidents WHERE lname LIKE ?",lName)
				if err != nil {
					panic(err.Error())
				}
				for row.Next(){
						var number int64
						var lname string
						var fname string
						var mInitial string
						var state string
						var party string
						var start int64
						var end int64
						var term string
						err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
						if err != nil {
							panic(err.Error())
						}
						USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
					}
				} else {
						state := r.FormValue("State")
						if state != "" {
						row, err := db.Query("SELECT * FROM presidents WHERE state=?",state)
						if err != nil {
							panic(err.Error())
						}
						for row.Next(){
							var number int64
							var lname string
							var fname string
							var mInitial string
							var state string
							var party string
							var start int64
							var end int64
							var term string
							err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
							if err != nil {
								panic(err.Error())
							}
						USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
					}
					}else {
						party := r.FormValue("Party")
						if party != "" {
							Party:="%"+party+"%"
							row, err := db.Query("SELECT * FROM presidents WHERE party=?",Party)
							if err != nil {
								panic(err.Error())
							}
							for row.Next(){
								var number int64
								var lname string
								var fname string
								var mInitial string
								var state string
								var party string
								var start int64
								var end int64
								var term string
								err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
								if err != nil {
									panic(err.Error())
								}
							USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
						}
						}else {
							start_Date := r.FormValue("Start")
							if start_Date != "" {
								int_start, _ := strconv.ParseInt(start_Date, 0, 64)
								row, err := db.Query("SELECT * FROM presidents WHERE start=?",int_start)
							if err != nil {
								panic(err.Error())
							}
							for row.Next(){
								var number int64
								var lname string
								var fname string
								var mInitial string
								var state string
								var party string
								var start int64
								var end int64
								var term string
								err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
								if err != nil {
									panic(err.Error())
								}
								USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
								
							}
							}else {
								end_Date := r.FormValue("End")
								if end_Date != "" {
									int_end, _ := strconv.ParseInt(end_Date, 0, 64)
									row, err := db.Query("SELECT * FROM presidents WHERE end=?",int_end)
									if err != nil {
										panic(err.Error())
									}
									for row.Next(){
										var number int64
										var lname string
										var fname string
										var mInitial string
										var state string
										var party string
										var start int64
										var end int64
										var term string
										err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
										if err != nil {
											panic(err.Error())
										}
										USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
										}
										}else {
											term_served := r.FormValue("Term")
											if term_served != "" {
												row, err := db.Query("SELECT * FROM presidents WHERE term=?",term_served)
												if err != nil {
													panic(err.Error())
												}
												for row.Next(){
													var number int64
													var lname string
													var fname string
													var mInitial string
													var state string
													var party string
													var start int64
													var end int64
													var term string
													err = row.Scan(&number, &lname, &fname, &mInitial, &state, &party, &start, &end, &term)
													if err != nil {
														panic(err.Error())
													}
													USPresidents=append(USPresidents, President{number, lname, fname, mInitial, state, party, start, end, term})
												}
											}
										}
									}
								}
							}
						}
					}
				}
		}
	tmpl.ExecuteTemplate(w, "listPresidentDB2019.html", USPresidents)
	defer db.Close()
}

func main() {
	//Replace 1111 with your localhost
	//Replace **********  with your webid
    log.Println("Server started on: http://localhost:8123")
    http.HandleFunc("/", ChoosePresident2019)
	http.HandleFunc("/listPresidentDB2019.html",ListPresident2019)
    http.Handle("/home/sshrest4/static/", http.StripPrefix("/home/sshrest4/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8123", nil)
}
