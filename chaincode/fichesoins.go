package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//FicheSoins declaration of struct
type FicheSoins struct {
	ObjectType            string
	UUID                  string
	IDContrat             string
	IDCompagnieAssurance  string
	IDHopital             string
	CodeAcheteurAssurance string
	DateDebut             string
	DateFin               string
	FicheSoinsPDF         string
	SignatureAcheteur     string
	SignatureCompagnie    string
	signatureHopital      string
}

func makeFicheSoinsFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) FicheSoins {
	ficheSoins := FicheSoins{}
	err := json.Unmarshal(bytes, &ficheSoins)
	panicErr(err)
	return ficheSoins
}

func makeBytesFromAsset(stub shim.ChaincodeStubInterface, ficheSoins FicheSoins) []byte {
	bytes, err := json.Marshal(ficheSoins)
	panicErr(err)
	return bytes
}

//CreateFicheSoinsOnLedger to create an FicheSoins on ledger
func CreateFicheSoinsOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string, iDContrat string,
	iDCompagnieAssurance string, iDHopital string, codeAcheteurAssurance string, dateDebut string, dateFin string,
	ficheSoinsPDF string, signatureAcheteur string, signatureCompagnie string, signatureHopital string) []byte {

	ficheSoins := FicheSoins{objectType, uuid, iDContrat, iDCompagnieAssurance, iDHopital, codeAcheteurAssurance,
		dateDebut, dateFin, ficheSoinsPDF, signatureAcheteur, signatureCompagnie, signatureHopital}

	ficheSoinsAsJSONBytes := makeBytesFromAsset(stub, ficheSoins)

	uuidIdexKeyFicheSoins := createIndexKey(stub, uuid, "fichesoins")

	putEntityInLedger(stub, uuidIdexKeyFicheSoins, ficheSoinsAsJSONBytes)
	return ficheSoinsAsJSONBytes
}

//CreateFicheSoins Core creation
func (t *FicheSoins) CreateFicheSoins(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	iDCompagnieAssurance := args[1]
	iDContrat := args[2]
	iDHopital := args[3]
	codeAcheteurAssurance := args[4]
	dateDebut := args[5]
	dateFin := args[6]
	ficheSoinsPDF := args[7]
	signatureAcheteur := args[8]
	signatureCompagnie := args[9]
	signatureHopital := args[10]

	uuidIndexKeyFicheSoins := createIndexKey(stub, uuid, "fichesoins")
	if checkEntityExist(stub, uuidIndexKeyFicheSoins) == true {
		return entityAlreadyExistMessage(stub, uuid, "fichesoins")
	}

	ficheSoins := CreateFicheSoinsOnLedger(stub, "fichesoins",
		uuid, iDCompagnieAssurance, iDContrat, iDHopital, codeAcheteurAssurance,
		dateDebut, dateFin, ficheSoinsPDF, signatureAcheteur, signatureCompagnie, signatureHopital)

	return succeed(stub, "FicheSoinsCreated", ficheSoins)
}

//GetFicheSoinsByID method to get an ficheSoins by id
func (t *FicheSoins) GetFicheSoinsByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetFicheSoinsByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "fichesoins")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "fichesoins")
	}
	ficheSoinsAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(ficheSoinsAsBytes)
}
