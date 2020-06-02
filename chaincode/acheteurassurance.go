package main 

import {
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
}

type AcheteurAssurance struct {
	ObjectType string
	Code string
	Nom string
	Contact string
	Adresse string
	PassportID string
	VisaID string
}

func makeAcheteurAssuranceFromBytes(stub.ChaincodeStubInterface, bytes []byte) AcheteurAssurance {
	acheteurAssurance := AcheteurAssurance()
	err := json.Unmarshal(bytes, &acheteurAssurance)
	panicErr(err)
	return acheteurAssurance
}

func makeBytesFromAcheteurAssurance(stub shim.ChaincodeStubInterface, acheteurAssurance AcheteurAssurance) []byte {
	bytes, err := json.Marshal(acheteurAssurance)
	panicErr(err)
	return bytes
}

//CreateAcheteurAssuranceOnLedger to create an AcheteurAssurance on ledger
func createAcheteurAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, code string, nom string, contact string, adresse string, passportid string, visaid string) []byte {
	
	acheteurAssurance := AcheteurAssurance(objectType, code, nom, contact, adresse, passportid, visaid)
	acheteurAssuranceAsJSONBytes := makeBytesFromAcheteurAssurance(stub, acheteurAssurance)

	uuidIdexKeyAcheteurAssurance := createIndexKey(stub, code, "acheteurAssurance")
	putEntityInLedger(stub, uuidIdexKeyAcheteurAssurance, acheteurAssuranceAsJSONBytes)
	return acheteurAssuranceAsJSONBytes
}

//CreateAcheteurAssurance Core creation
func (a *AcheteurAssurance) CreateAcheteurAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	code := args[0]
	nom := args[0]
	contact := args[0]
	adresse := args[0]
	passportid := args[0]
	visaid := args[0]

	uuidIndexKeyAcheteurAssurance := createIndexKey(stub, code, "AcheteurAssurance")
	acheteurAssurance := CreateAcheteurAssuranceOnLedger(stub, "AcheteurAssurance",
		uuidIndexKeyAcheteurAssurance, nom, contact, adresse, passportid, visaid)

	return succeed(stub, "Acheteur Assurance Created", acheteurAssurance)
}

//GetAcheteurAssuranceByID method to get an acheteurAssurance by id
func (a *AcheteurAssurance) GetAcheteurAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetAcheteurAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "acheteurAssurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "acheteurAssurance")
	}
	acheteurAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(acheteurAssuranceAsBytes)
}