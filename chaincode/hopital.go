package main 

import {
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
}

type Hopital struct {
	ObjectType string
	Code string
	Nom string
	Contact string
	Adresse string
}

func makeHopitalFromBytes(stub.ChaincodeStubInterface, bytes []byte) Hopital {
	hopital := Hopital()
	err := json.Unmarshal(bytes, &hopital)
	panicErr(err)
	return hopital
}

func makeBytesFromHopital(stub shim.ChaincodeStubInterface, hopital Hopital) []byte {
	bytes, err := json.Marshal(hopital)
	panicErr(err)
	return bytes
}

//CreateHopitalOnLedger to create an Hopital on ledger
func createHopitalOnLedger(stub shim.ChaincodeStubInterface, objectType string, code string, nom string, contact string, adresse string) {
	
	hopital := Hopital(objectType, code, nom, contact, adresse)
	hopitalAsJSONBytes := makeBytesFromHopital(stub, hopital)

	uuidIdexKeyHopital := createIndexKey(stub, code, "hopital")
	putEntityInLedger(stub, uuidIdexKeyHopital, hopitalAsJSONBytes)
}

//CreateHopital Core creation
func (h *Hopital) CreateHopital(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	code := args[0]
	nom := args[0]
	contact := args[0]
	adresse := args[0]

	uuidIndexKeyHopital := createIndexKey(stub, code, "Hopital")
	hopital := CreateHopitalOnLedger(stub, "Hopital", uuidIndexKeyHopital, nom, contact, adresse)

	return succeed(stub, "HopitalCreated", hopital)
}

//GetHopitalByID method to get an hopital by id
func (h *Hopital) GetHopitalByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetHopitalByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "hopital")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "hopital")
	}
	hopitalAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(hopitalAsBytes)
}

