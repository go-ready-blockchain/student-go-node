# STUDENT NODE

## Blockchain Implementation in GoLang For Placement System

## The Consensus Algorithm implemented in Blockchain System is a combination of Proof Of Work and Proof Of Elapsed Time


### Run `go run main.go` to Start the Server and listen on localhost:8081

### Usage :

#### To Print Usage
####    Make POST request to `/usage`

#### To Advance the Pipeline - 

#### To Add a New Student
####    Make POST request to `/student` with body -
```json
{
    "Usn": "1MS16CS034",
    "Branch": "CSE",
    "Name": "Gaurav",
    "Gender": "Male",
    "Dob": "30-10-1998",
    "Cgpa": "9",
    "Perc10th": "90",
    "Perc12th": "90",
    "Backlog": false,
    "Email": "gauravkarkal@gmail.com",
    "Mobile": "8867454545",
    "Staroffer": true
}
```

#### Part of the Pipeline -  

#### To Handle Request and Initiate Creation of Request Block
####    Make GET request to `/handlerequest` with Query Params -
```json
Key :   Value

approval: true
company: JPMC
name: 1MS16CS034

```

#### Testing

#### Test Direct Request to Student
####    Make POST request to `/request-student` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```



