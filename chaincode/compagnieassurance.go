package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CompagnieAssurance declaration of the struct
type CompagnieAssurance struct {
	ObjectType string
	Code       string
	Nom        string
	Contact    string
	Adresse    string
}

func makeCompagnieAssuranceFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) CompagnieAssurance {
	compagnieAssurance := CompagnieAssurance{}
	err := json.Unmarshal(bytes, &compagnieAssurance)
	panicErr(err)
	return compagnieAssurance
}

func makeBytesFromOrganization(stub shim.ChaincodeStubInterface, compagnieAssurance CompagnieAssurance) []byte {
	bytes, err := json.Marshal(compagnieAssurance)
	panicErr(err)
	return bytes
}

//CreateCompagnieAssuranceOnLedger to create an CompagnieAssurance on ledger
func CreateCompagnieAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, code string,
	nom string, contact string, adresse string) []byte {

	compagnieAssurance := CompagnieAssurance{objectType, code, nom, contact, adresse}
	compagnieAssuranceAsJSONBytes := makeBytesFromOrganization(stub, compagnieAssurance)

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, code, "organization")
	putEntityInLedger(stub, uuidIndexKeyCompagnieAssurance, compagnieAssuranceAsJSONBytes)
	return compagnieAssuranceAsJSONBytes

}

//CreateCompagnieAssurance Core creation
func (t *CompagnieAssurance) CreateCompagnieAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	code := args[0]
	nom := args[0]
	contact := args[0]
	adresse := args[0]

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, code, "CompagnieAssurance")
	compagnieAssurance := CreateCompagnieAssuranceOnLedger(stub, "CompagnieAssurance",
		uuidIndexKeyCompagnieAssurance, nom, contact, adresse)

	return succeed(stub, "CompagnieAssuranceCreated", compagnieAssurance)
}

//GetCompagnieAssuranceByID method to get an organization by id
func (t *CompagnieAssurance) GetCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetCompagnieAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "organization")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "organization")
	}
	organizationAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(organizationAsBytes)
}