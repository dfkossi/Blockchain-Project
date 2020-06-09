package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//ContratAssurance declaration of the struct
type ContratAssurance struct {
	ObjectType            string
	UUID                  string
	IDCompagnieAssurance  CompagnieAssurance
	CodeAcheteurAssurance string
	DateDebut             string
	DateFin               string
	ContratAssurancePDF   string
	SignatureAcheteur     string
	SignatureCompagnie    string
}

func makeContratAssuranceFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) ContratAssurance {
	contratAssurance := ContratAssurance{}
	err := json.Unmarshal(bytes, &contratAssurance)
	panicErr(err)
	return contratAssurance
}

func makeBytesFromContratAssurance(stub shim.ChaincodeStubInterface, contratAssurance ContratAssurance) []byte {
	bytes, err := json.Marshal(contratAssurance)
	panicErr(err)
	return bytes
}

//CreateContratAssuranceOnLedger to create an CompagnieAssurance on ledger
func CreateContratAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	iDCompagnieAssurance CompagnieAssurance, codeAcheteurAssurance string, dateDebut string,
	dateFin string, contratAssurancePDF string, signatureAcheteur string, signatureCompagnie string) []byte {

	contratAssurance := ContratAssurance{objectType, uuid, iDCompagnieAssurance, codeAcheteurAssurance,
		dateDebut, dateFin, contratAssurancePDF, signatureAcheteur, signatureCompagnie}
	contratAssuranceAsJSONBytes := makeBytesFromContratAssurance(stub, contratAssurance)

	uuidIndexKeyContratAssurance := createIndexKey(stub, uuid, "contratassurance")

	putEntityInLedger(stub, uuidIndexKeyContratAssurance, contratAssuranceAsJSONBytes)
	return contratAssuranceAsJSONBytes
}

//CreateContratAssurance Core creation
func (t *ContratAssurance) CreateContratAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	iDCompagnieAssurance := args[1]
	codeAcheteurAssurance := args[2]
	dateDebut := args[3]
	dateFin := args[4]
	contratAssurancePDF := args[5]
	signatureAcheteur := args[6]
	signatureCompagnie := args[7]

	uuidIndexKeyContratAssurance := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKeyContratAssurance) == true {
		return entityAlreadyExistMessage(stub, uuid, "contratassurance")
	}

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, iDCompagnieAssurance, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKeyCompagnieAssurance) == true {
		return entityAlreadyExistMessage(stub, iDCompagnieAssurance, "compagnieassurance")
	}
	compagnieAssurance := getEntityFromLedger(stub, uuidIndexKeyCompagnieAssurance)
	compagnieAssuranceAsJSONBytes := makeCompagnieAssuranceFromBytes(stub, compagnieAssurance)

	contratAssurance := CreateContratAssuranceOnLedger(stub, "contratassurance",
		uuid, compagnieAssuranceAsJSONBytes, codeAcheteurAssurance, dateDebut, dateFin,
		contratAssurancePDF, signatureAcheteur, signatureCompagnie)

	return succeed(stub, "ContratAssuranceCreated", contratAssurance)
}

//GetContratAssuranceByID method to get an contratAssurance by id
func (t *ContratAssurance) GetContratAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetContratAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}
	contratAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(contratAssuranceAsBytes)
}

//UpdateCompagnieAssuranceByID method to update an compagnieassurance by id
func (t *ContratAssurance) UpdateCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UpdateCompagnieAssuranceByID - Start")

	uuid := args[0]
	newNom := args[1]
	newContact := args[2]
	newAdresse := args[3]

	uuidIndexKey := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}
	organizationAsBytes := getEntityFromLedger(stub, uuidIndexKey)
	contratassurance := makeCompagnieAssuranceFromBytes(stub, organizationAsBytes)

	contratassurance.Nom = newNom
	contratassurance.Contact = newContact
	contratassurance.Adresse = newAdresse

	compagnieAssuranceAsJSONBytes := makeBytesFromCompagnieAssurance(stub, contratassurance)

	putEntityInLedger(stub, uuidIndexKey, compagnieAssuranceAsJSONBytes)
	return succeed(stub, "CompagnieAssuranceUpdatedEvent", compagnieAssuranceAsJSONBytes)

}
