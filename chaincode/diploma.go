package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"

	// "strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

type DiplomaPublicDetails struct {
	UniversityId   string `json:"universityId"`
	UniversityName string `json:"universityName"`
	DiplomaId      string `json:"diplomaId"`
	StudentName    string `json:"studentName"`
	CourseName     string `json:"courseName"`
	EndDate        string `json:"endDate"`
	StudentCpf     string `json:"studentCpf"`
}

// type DiplomaPrivateDetails struct {
// 	DiplomaId 			string `json:"diplomaId"`
// 	DiplomaHash 		string `json:"diplomaHash"`
// 	Course 				Course
// 	Student 			Student
// }

// type Course struct{
// 	CourseName			string `json:"courseName"`
// 	BeginDate			string `json:"beginDate"`
// 	EndDate				string `json:"endDate"`
// 	Teachers 			Teachers
// }

// type Teachers struct {
// 	TeacherName			string `json:"teacherName"`
// 	Subject       		string `json:"subject"`
// }

// type Student struct {
// 	StudentId 			string `json:"studentId"`
// 	StudentName 		string `json:"studentName"`
// 	StudentEmail 		string `json:"studentEmail"`
// 	StudentCpf 			string `json:"studentCpf"`
// }

// type DiplomaPublicDetails struct {

// 	Universidade struct {
// 		UniversityId   string `json:"universityId"`
// 		UniversityName string `json:"universityName"`
// 	} `json:"universidade"`
// 	Diploma struct {
// 		DiplomaID   string `json:"diplomaId"`
// 		DiplomaHash string `json:"-"`
// 		Curso       struct {
// 			CourseName  string `json:"courseName"`
// 			BeginDate   string `json:"-"`
// 			EndDate     string `json:"endDate"`
// 			Professores []struct {
// 				TeacherName string `json:"-"`
// 				Subject     string `json:"-"`
// 			} `json:"professores"`
// 		} `json:"curso"`
// 		Aluno struct {
// 			StudentID    string `json:"-"`
// 			StudentEmail string `json:"-"`
// 			StudentName  string `json:"studentName"`
// 			StudentCpf   string `json:"studentCpf"`
// 		} `json:"aluno"`
// 	} `json:"diploma"`
// }

type DiplomaPrivateDetails struct {
	Collection   string `json:"collection"`
	Universidade struct {
		UniversityId   string `json:"universityId"`
		UniversityName string `json:"universityName"`
	} `json:"universidade"`
	Diploma struct {
		DiplomaID   string `json:"diplomaId"`
		DiplomaHash string `json:"diplomaHash"`
		Curso       struct {
			CourseName  string `json:"courseName"`
			BeginDate   string `json:"beginDate"`
			EndDate     string `json:"endDate"`
			Professores []struct {
				TeacherName string `json:"teacherName"`
				Subject     string `json:"subject"`
			} `json:"professores"`
		} `json:"curso"`
		Aluno struct {
			StudentID    string `json:"studentId"`
			StudentEmail string `json:"studentEmail"`
			StudentName  string `json:"studentName"`
			StudentCpf   string `json:"studentCpf"`
		} `json:"aluno"`
	} `json:"diploma"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryDiploma" {
		return s.queryDiploma(APIstub, args)
	} else if function == "criarDiploma" {
		return s.criarDiploma(APIstub, args)
	} else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	}
	// else if function == "searchDiplomas" {
	// 	return s.searchDiplomas(APIstub)
	// }

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) criarDiploma(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	var err error
	var jsonpayload = args[0]
	fmt.Println("jsonpayload ", jsonpayload)
	bytes := []byte(args[0])
	var collection = args[1]
	var key = args[2]

	var diplomaPrivateDetails DiplomaPrivateDetails

	if err := json.Unmarshal(bytes, &diplomaPrivateDetails); err != nil {
		return shim.Error(err.Error())
	}
	err = APIstub.PutPrivateData(collection, key, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// var dat map[string]interface{}
	// if err := json.Unmarshal(bytes, &dat); err != nil {
	//     panic(err)
	// }
	// fmt.Println("datt ", dat)

	// res := []DiplomaPublicDetails{}

	// json.Unmarshal(bytes, &res)

	// fmt.Println("DiplomaPublicDetails ", res)

	diplomaPublicDetails := &DiplomaPublicDetails{
		UniversityId:   args[3],
		UniversityName: args[4],
		DiplomaId:      args[5],
		StudentName:    args[6],
		CourseName:     args[7],
		EndDate:        args[8],
		StudentCpf:     args[9],
	}

	diplomaPublicDetailsJSONasBytes, err := json.Marshal(diplomaPublicDetails)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = APIstub.PutPrivateData("collectionDiploma", key, diplomaPublicDetailsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// diplomaAsBytes, _ := json.Marshal(res)
	// APIstub.PutPrivateData(collection, key, diplomaAsBytes)

	// json.Unmarshal(bytes, &diplomaPublicDetails)
	// fmt.Println(diplomaPublicDetails.Universidade.UniversityID)

	// diplomaPublicDetailsAsBytes, _ := json.Marshal(diplomaPublicDetails)
	// fmt.Println(diplomaPublicDetails.Universidade.UniversityID)

	// APIstub.PutPrivateData("collectionDiploma", key, diplomaPublicDetailsAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) searchForDiploma(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	diploma := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"studentName\":\"%s\"}}", diploma)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryDiploma(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var searchKey, jsonResp, searchCollection string
	fmt.Println("arguments ", args)
	var err error

	searchKey = args[0]
	searchCollection = args[1]
	fmt.Println("searchKey ", searchKey)

	valAsbytes, err := APIstub.GetPrivateData(searchCollection, searchKey) //get the diploma private details from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get private details for " + searchKey + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"diploma private details does not exist: " + searchKey + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := APIstub.GetPrivateDataQueryResult("collectionDiploma", queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// func (s *SmartContract) searchDiplomas(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	var searchKey = args[0]

// 	resultsIterator, err := APIstub.GetPrivateDataQueryResult("collectiondiplomas", searchKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing QueryRecords
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(queryResponse.Key)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.WriteString(string(queryResponse.Value))
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

// 	return buffer.Bytes(), nil
// }

func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
