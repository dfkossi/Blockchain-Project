package main

import (
	"encoding/json"
	"fmt"

	//_ "utils.go"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CompagnieAssurance declaration of the struct
type CompagnieAssurance struct {
	ObjectType string
	UUID       string
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

func makeBytesFromCompagnieAssurance(stub shim.ChaincodeStubInterface, compagnieAssurance CompagnieAssurance) []byte {
	bytes, err := json.Marshal(compagnieAssurance)
	panicErr(err)
	return bytes
}

//CreateCompagnieAssuranceOnLedger to create an CompagnieAssurance on ledger
func CreateCompagnieAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	nom string, contact string, adresse string) []byte {

	compagnieAssurance := CompagnieAssurance{objectType, uuid, nom, contact, adresse}
	compagnieAssuranceAsJSONBytes := makeBytesFromCompagnieAssurance(stub, compagnieAssurance)

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, uuid, "compagnieassurance")

	putEntityInLedger(stub, uuidIndexKeyCompagnieAssurance, compagnieAssuranceAsJSONBytes)
	return compagnieAssuranceAsJSONBytes

}

//CreateCompagnieAssurance Core creation
func (t *CompagnieAssurance) CreateCompagnieAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	nom := args[1]
	contact := args[2]
	adresse := args[3]

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKeyCompagnieAssurance) == true {
		return entityAlreadyExistMessage(stub, uuid, "compagnieassurance")
	}

	compagnieAssurance := CreateCompagnieAssuranceOnLedger(stub, "compagnieassurance",
		uuid, nom, contact, adresse)
	return succeed(stub, "CompagnieAssuranceCreated", compagnieAssurance)
}

//GetCompagnieAssuranceByID method to get an compagnieAssurance by id
func (t *CompagnieAssurance) GetCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetCompagnieAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "compagnieassurance")
	}
	compagnieAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(compagnieAssuranceAsBytes)
}
