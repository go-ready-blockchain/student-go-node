package Init

import (
	
	"fmt"
	

	"github.com/jugalw13/student-go-node/blockchain"
	"github.com/jugalw13/student-go-node/security"
	"github.com/jugalw13/student-go-node/student"
	"github.com/jugalw13/student-go-node/utils"
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
	fmt.Println(stud)

	utils.StoreStudentData(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)

	//StoreStudentDataInDb(student.EncodeToBytes(stud))

}

