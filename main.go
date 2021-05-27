package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"time"
)

type Student struct {
	SID     uuid.UUID  `json:"id"`
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Class   string     `json:"class"`
	Subject string     `json:"subject"`
}

var Students []Student

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if Students != nil {
		json.NewEncoder(w).Encode(Students)
	} else {
		json.NewEncoder(w).Encode("No students")
	}
}

func getSingleStudent(w http.ResponseWriter, r *http.Request) {
	stud_name := mux.Vars(r)["name"]
	flag := 0
	stud := Student{}
	for _, student := range Students {
		if student.Name == stud_name {
			flag = 1
			stud = student
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if flag == 1 {
		json.NewEncoder(w).Encode(stud)
	} else {
		json.NewEncoder(w).Encode("Student Not Found")
	}
}

func delStudent(w http.ResponseWriter, r *http.Request) {
	stud_name := mux.Vars(r)["name"]
	flag := 0
	stud := Student{}
	for index, student := range Students {
		if student.Name == stud_name {
			flag = 1
			stud = student
			Students = append(Students[:index], Students[index+1:]...)
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	if flag == 1 {
		json.NewEncoder(w).Encode(stud.Name+" Deleted")
	} else {
		json.NewEncoder(w).Encode("Student Not Found")
	}
}

func addNewStudent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student Student
	json.Unmarshal(reqBody, &student)
	Students = append(Students, student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func handleAllRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students", addNewStudent).Methods("POST")
	router.HandleFunc("/students/{name}", getSingleStudent).Methods("GET")
	router.HandleFunc("/students/{name}", delStudent).Methods("DELETE")
	
	srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("Server UP on port 8000")
	Students = []Student{
		{
			uuid.New(),"ABC",18,"C1","M1",
		},
		{
			uuid.New(),"XYZ",19,"C2","M2",
		},
	}
	handleAllRequests()
}
