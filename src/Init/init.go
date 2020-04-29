package Init

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/jugalw13/student-go-node/blockchain"
	"github.com/jugalw13/student-go-node/security"
	"github.com/jugalw13/student-go-node/student"
)

func InitializeBlockChain() {
	blockchain.InitBlockChain()
	InitNodes()
}

func InitNodes() {

	security.GenerateAcademicDeptKeys()

	security.GeneratePlacementDeptKeys()

}
func InitCompanyNode(company string) {
	security.GenerateCompanyKeys(company)
}

func InitStudentNode(usn string, branch string, name string, gender string, dob string, perc10th string, perc12th string, cgpa string, backlog bool, email string, mobile string, staroffer bool) {

	security.GenerateStudentKeys(usn)

	stud := student.EnterStudentData(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)

	StoreStudentDataInDb(student.EncodeToBytes(stud))

}

func StoreStudentDataInDb(jsonbytes []byte) bool {

	url := "https://hn86a2dvf0.execute-api.us-east-1.amazonaws.com/default/blockchaindb"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonbytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		fmt.Println("Student successfully added to DB!")
		return true
	}

	fmt.Println("Student Failed to add to DB!")
	return false

}
