package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/symon-nascimento/api/model"
	"github.com/symon-nascimento/api/worker"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func WorkerPrint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, worker.Print())
}

func WorkerSync(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, worker.Sync())
}

func WorkerPersist(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, worker.Persist())
}

func (h handler) AddStudent(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var student model.Student
	json.Unmarshal(body, &student)

	if result := h.DB.Create(&student); result.Error != nil {
		fmt.Println(result.Error)
	}
	// Imprimindo o valor do body
	fmt.Print("User created: ", student.Name)
	// Devolvendo o req
	json.NewEncoder(w).Encode(student.Name)
}

func (h handler) RmStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Print(w, vars["id"])
	id, _ := strconv.Atoi(vars["id"])

	var student model.Student

	if result := h.DB.First(&student, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode("Deleted")

	// Deletando usuário
	h.DB.Delete(&student)
}

func (h handler) UpStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Print(w, vars["id"])

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateStudent model.Student
	json.Unmarshal(body, &updateStudent)

	var student model.Student

	if result := h.DB.First(&student, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	student.Name = updateStudent.Name
	h.DB.Save(&student)

	// Imprimindo o valor do body
	fmt.Print("User update: ", student.Name)
	// Devolvendo o req
	json.NewEncoder(w).Encode(student.Name)
}

func (h handler) FindOneStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])
	fmt.Print("ID ->",  id)

	var student model.Student

	if result := h.DB.First(&student, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	query:= h.DB.First(&student)
	fmt.Print(query)

	json.NewEncoder(w).Encode(student)
}