package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type President struct {
	Num      int64
	LName    string
	FName    string
	MInitial string
	State    string
	Party    string
	Start    int64
	End      int64
	Term     string
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func ChoosePresident2019(w http.ResponseWriter, r *http.Request) {
	WW2Presidents := []President{}
	WW2Presidents = append(WW2Presidents, President{Num: 34, LName: "Eisenhower", FName: "Dwight", MInitial: " D. ", State: "Kansas", Party: "Republican", Start: 1953, End: 1961, Term: "42 & 43"})
	WW2Presidents = append(WW2Presidents, President{Num: 35, LName: "Kennedy", FName: "John", MInitial: " F. ", State: "Massachusetts", Party: "Democrat", Start: 1961, End: 1963, Term: "44"})
	WW2Presidents = append(WW2Presidents, President{Num: 36, LName: "Johnson", FName: "Lyndon", MInitial: " B. ", State: "Texas", Party: "Democrat", Start: 1963, End: 1969, Term: "44 & 45"})
	WW2Presidents = append(WW2Presidents, President{Num: 37, LName: "Nixon", FName: "Richard", MInitial: " M. ", State: "California", Party: "Republican", Start: 1969, End: 1974, Term: "46 & 47"})
	WW2Presidents = append(WW2Presidents, President{Num: 38, LName: "Ford", FName: "Gerald", MInitial: " ", State: "Michigan", Party: "Republican", Start: 1974, End: 1977, Term: "47"})
	WW2Presidents = append(WW2Presidents, President{Num: 39, LName: "Carter", FName: "Jimmy", MInitial: " ", State: "Georgia", Party: "Democrat", Start: 1977, End: 1981, Term: "48"})
	WW2Presidents = append(WW2Presidents, President{Num: 40, LName: "Reagan", FName: "Ronald", MInitial: " ", State: "California", Party: "Republican", Start: 1981, End: 1989, Term: "49 & 50"})
	WW2Presidents = append(WW2Presidents, President{Num: 41, LName: "Bush", FName: "George", MInitial: " H. ", State: "Texas", Party: "Republican", Start: 1989, End: 1993, Term: "51"})
	WW2Presidents = append(WW2Presidents, President{Num: 42, LName: "Clinton", FName: "Bill", MInitial: " ", State: "Arkansas", Party: "Democrat", Start: 1993, End: 2001, Term: "52 & 53"})
	WW2Presidents = append(WW2Presidents, President{Num: 43, LName: "Bush", FName: "George", MInitial: " W. ", State: "Texas", Party: "Republican", Start: 2001, End: 2009, Term: "54 & 55"})
	WW2Presidents = append(WW2Presidents, President{Num: 44, LName: "Obama", FName: "Barack", MInitial: " H. ", State: "Illionie", Party: "Democrat", Start: 2009, End: 2017, Term: "56 & 57"})
	WW2Presidents = append(WW2Presidents, President{Num: 45, LName: "Trump", FName: "Donald", MInitial: " J. ", State: "New York", Party: "Republican", Start: 2017, End: 2019, Term: "58"})
	tmpl.ExecuteTemplate(w, "choosePresident2019.html", WW2Presidents)
}

func ListPresident2019(w http.ResponseWriter, r *http.Request) {
	WW2Presidents := []President{}
	WW2Presidents = append(WW2Presidents, President{Num: 34, LName: "Eisenhower", FName: "Dwight", MInitial: " D. ", State: "Kansas", Party: "Republican", Start: 1953, End: 1961, Term: "42 & 43"})
	WW2Presidents = append(WW2Presidents, President{Num: 35, LName: "Kennedy", FName: "John", MInitial: " F. ", State: "Massachusetts", Party: "Democrat", Start: 1961, End: 1963, Term: "44"})
	WW2Presidents = append(WW2Presidents, President{Num: 36, LName: "Johnson", FName: "Lyndon", MInitial: " B. ", State: "Texas", Party: "Democrat", Start: 1963, End: 1969, Term: "44 & 45"})
	WW2Presidents = append(WW2Presidents, President{Num: 37, LName: "Nixon", FName: "Richard", MInitial: " M. ", State: "California", Party: "Republican", Start: 1969, End: 1974, Term: "46 & 47"})
	WW2Presidents = append(WW2Presidents, President{Num: 38, LName: "Ford", FName: "Gerald", MInitial: " ", State: "Michigan", Party: "Republican", Start: 1974, End: 1977, Term: "47"})
	WW2Presidents = append(WW2Presidents, President{Num: 39, LName: "Carter", FName: "Jimmy", MInitial: " ", State: "Georgia", Party: "Democrat", Start: 1977, End: 1981, Term: "48"})
	WW2Presidents = append(WW2Presidents, President{Num: 40, LName: "Reagan", FName: "Ronald", MInitial: " ", State: "California", Party: "Republican", Start: 1981, End: 1989, Term: "49 & 50"})
	WW2Presidents = append(WW2Presidents, President{Num: 41, LName: "Bush", FName: "George", MInitial: " H. ", State: "Texas", Party: "Republican", Start: 1989, End: 1993, Term: "51"})
	WW2Presidents = append(WW2Presidents, President{Num: 42, LName: "Clinton", FName: "Bill", MInitial: " ", State: "Arkansas", Party: "Democrat", Start: 1993, End: 2001, Term: "52 & 53"})
	WW2Presidents = append(WW2Presidents, President{Num: 43, LName: "Bush", FName: "George", MInitial: " W. ", State: "Texas", Party: "Republican", Start: 2001, End: 2009, Term: "54 & 55"})
	WW2Presidents = append(WW2Presidents, President{Num: 44, LName: "Obama", FName: "Barack", MInitial: " H. ", State: "Illionie", Party: "Democrat", Start: 2009, End: 2017, Term: "56 & 57"})
	WW2Presidents = append(WW2Presidents, President{Num: 45, LName: "Trump", FName: "Donald", MInitial: " J. ", State: "New York", Party: "Republican", Start: 2017, End: 2019, Term: "58"})
	FoundPresidents := []President{}
	if r.Method == "POST" {
		president_id := r.FormValue("ID")
		if president_id != "" {
			int_ID, _ := strconv.ParseInt(president_id, 0, 64)
			for _, value := range WW2Presidents {
				if int_ID == value.Num {
					FoundPresidents = append(FoundPresidents, value)
				}
			}
		}else {
			first_name := r.FormValue("FName")
			if first_name != "" {
				for _, value := range WW2Presidents {
					if strings.Contains(strings.TrimSpace(strings.ToLower(value.FName)),strings.TrimSpace(strings.ToLower(first_name))) {
						FoundPresidents = append(FoundPresidents, value)
					}
				}
			}else {
				last_name := r.FormValue("LName")
				if last_name != "" {
					for _, value := range WW2Presidents {
						if strings.Contains(strings.TrimSpace(strings.ToLower(value.LName)),strings.TrimSpace(strings.ToLower(last_name))) {
							FoundPresidents = append(FoundPresidents, value)
						}
					}
				} else {
					state := r.FormValue("State")
					if state != "" {
						for _, value := range WW2Presidents {
							if strings.Contains(strings.TrimSpace(strings.ToLower(value.State)),strings.TrimSpace(strings.ToLower(state))) {
								FoundPresidents = append(FoundPresidents, value)
							}
						}
					}else {
						party := r.FormValue("Party")
						if party != "" {
							for _, value := range WW2Presidents {
								if strings.Contains(strings.TrimSpace(strings.ToLower(value.Party)),strings.TrimSpace(strings.ToLower(party))) {
									FoundPresidents = append(FoundPresidents, value)
								}
							}
						} else {
							start_Date := r.FormValue("Start")
							if start_Date != "" {
								int_start, _ := strconv.ParseInt(start_Date, 0, 64)
								for _, value := range WW2Presidents {
									if value.Start == int_start {
										FoundPresidents = append(FoundPresidents, value)
									}
								}
							} else {
								end_Date := r.FormValue("End")
								if end_Date != "" {
									int_end, _ := strconv.ParseInt(end_Date, 0, 64)
									for _, value := range WW2Presidents {
										if value.End == int_end {
											FoundPresidents = append(FoundPresidents, value)
										}
									}
								} else {
									term_served := r.FormValue("Term")
									if term_served != "" {
										for _, value := range WW2Presidents {
											if strings.Contains(strings.TrimSpace(strings.ToLower(value.Term)), strings.TrimSpace(strings.ToLower(term_served))) {
												FoundPresidents = append(FoundPresidents, value)
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
	}

	tmpl.ExecuteTemplate(w, "listPresident2019.html", FoundPresidents)
}

func main() {
	log.Println("Server started on: http://localhost:8123")
	http.HandleFunc("/", ChoosePresident2019)
	http.HandleFunc("/listPresident2019.html", ListPresident2019)
	http.Handle("/home/sshrest4/static/", http.StripPrefix("/home/sshrest4/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8123", nil)
}
