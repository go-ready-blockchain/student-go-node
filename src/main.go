package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-ready-blockchain/blockchain-go-core/Init"
	"github.com/go-ready-blockchain/blockchain-go-core/blockchain"
	"github.com/go-ready-blockchain/blockchain-go-core/logger"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("Make POST request to /student \tTo Add a New Student")
	fmt.Println("Make GET request to /handlerequest \tHandle Request and Initiate Creation of Request Block")
	fmt.Println("Make POST request to /request-student \t Test Direct Request to Student")
}

func addStudent(usn string, branch string, name string, gender string, dob string, perc10th float32, perc12th float32, cgpa float32, backlog bool, email string, mobile string, staroffer bool) {
	fmt.Println("\nInitializing new Student\n")
	Init.InitStudentNode(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)
	fmt.Println("Student Added!")

}

func calladdStudent(w http.ResponseWriter, r *http.Request) {
	name := time.Now().String()
	logger.FileName = "Add Student" + name + ".log"
	logger.NodeName = "Student Node"
	logger.CreateFile()

	type jsonBody struct {
		Usn       string  `json:"Usn"`
		Branch    string  `json:"Branch"`
		Name      string  `json:"Name"`
		Gender    string  `json:"Gender"`
		Dob       string  `json:"Dob"`
		Perc10th  float32 `json:"Perc10th"`
		Perc12th  float32 `json:"Perc12th"`
		Cgpa      float32 `json:"Cgpa"`
		Backlog   bool    `json:"Backlog"`
		Email     string  `json:"Email"`
		Mobile    string  `json:"Mobile"`
		StarOffer bool    `json:"StarOffer"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	logger.UploadToS3Bucket(logger.NodeName)

	logger.DeleteFile()

	addStudent(b.Usn, b.Branch, b.Name, b.Gender, b.Dob, b.Perc10th, b.Perc12th, b.Cgpa, b.Backlog, b.Email, b.Mobile, b.StarOffer)

	message := "Student Added!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func requestBlock(name string, company string) {
	fmt.Println("\nCreating the requested block\n")
	blockchain.InitBlockInBuffer(name, company)
	fmt.Println("Requested Block Initialized!")
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	name := time.Now().String()
	logger.FileName = "Handle Student Request" + name + ".log"
	logger.NodeName = "Student Node"
	logger.CreateFile()

	type jsonBody struct {
		Approval bool   `json:"approval"`
		Name     string `json:"name"`
		Company  string `json:"company"`
	}

	decoder := r.URL.Query()
	approval := decoder["approval"][0]
	Approval, _ := strconv.ParseBool(approval)
	Company := decoder["company"][0]
	Name := decoder["name"][0]
	fmt.Println(Approval, Company)

	if !Approval {
		logger.UploadToS3Bucket(logger.NodeName)

		logger.DeleteFile()
		fmt.Println("Student :", Name, "Rejected Request for Data for Company: ", Company)
		w.Write([]byte(string("Student : " + Name + " Rejected Request for Data for Company: " + Company)))
		return
	}
	requestBlock(Name, Company)

	logger.UploadToS3Bucket(logger.NodeName)

	logger.DeleteFile()

	message := "Requested Block Initialized!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))

	fmt.Println("\n\nSending Notification to Academic Dept for Verification\n\n")
	callAcademicDeptVerification(Name, Company)

}

func request_student(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	requestBlock(b.Name, b.Company)

	message := "Requested Block Initialized!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))

	fmt.Println("\n\nSending Notification to Academic Dept for Verification\n\n")
	callAcademicDeptVerification(b.Name, b.Company)
}

func callAcademicDeptVerification(name string, company string) {
	reqBody, err := json.Marshal(map[string]string{
		"name":    name,
		"company": company,
	})
	if err != nil {
		print(err)
	}
	resp, err := http.Post("http://localhost:8083/verify-AcademicDept",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func callprintUsage(w http.ResponseWriter, r *http.Request) {

	printUsage()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Usage!!"
	w.Write([]byte(message))
}

func main() {
	port := "8081"
	http.HandleFunc("/student", calladdStudent)
	http.HandleFunc("/handlerequest", handlerequest)
	http.HandleFunc("/request-student", request_student)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("Server listening on localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
