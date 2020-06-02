package main 

import {
	"encoding/json"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
}

type ContratAssurance struct {
	IdContrat string
	IdCompagnieAssurance string
	CodeAcheteurAssurance string
	DateDebut datetime
	DateFin datetime
	FichierContrat string
	SignatureAcheteur string
	SignatureCompagnie string
}

func makeContratAssuranceFromBytes(stub.ChaincodeStubInterface, bytes []byte) ContratAssurance {
	asset := ContratAssurance()
	err := json.Unmarshal(bytes, &contratAssurance)
	panicErr(err)
	return contratAssurance
}

func makeBytesFromAsset(stub shim.ChaincodeStubInterface, contratAssurance ContratAssurance) []byte {
	bytes, err := json.Marshal(contratAssurance)
	panicErr(err)
	return bytes
}

func createAssetOnLedger(stub shim.ChaincodeStubInterface, objectType string, idContrat string, idCompagnieAssurance string, codeAcheteurAssurance string, dateDebut datetime, dateFin datetime, fichierContrat string, signatureAcheteur string, signatureCompagnie string) {
	
	contratAssurance := ContratAssurance(objectType, idContrat, idCompagnieAssurance, codeAcheteurAssurance, dateDebut, dateFin, fichierContrat, signatureAcheteur, signatureCompagnie)
	contratAssuranceAsJSONBytes := makeBytesFromAsset(stub, contratAssurance)

	uuidIdexKeyContratAssurance := createIndexKey(stub, idContrat, "asset")
	//putEntityInLedger(stub, uuidIdexKeyContratAssurance, contratAssuranceAsJSONBytes)
	return contratAssuranceAsJSONBytes
}

//CreateContratAssurance Core creation
func (c *ContratAssurance) CreateContactAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	idContrat := args[0]
	idCompagnieAssurance := args[0]
	codeAcheteurAssurance := args[0]
	dateDebut := args[0]
	dateFin := args[0]
	fichierContrat := args[0]
	signatureAcheteur := args[0]
	signatureCompagnie:= args[0]

	uuidIndexKeyContratAssurance := createIndexKey(stub, idContrat, "ContratAssurance")
	contratAssurance := CreateContratAssuranceOnLedger(stub, "ContratAssurance",
		uuidIndexKeyContratAssurance, idCompagnieAssurance, codeAcheteurAssurance, dateDebut, dateFin, fichierContrat, signatureAcheteur, signatureCompagnie)

	return succeed(stub, "Contrat Assurance Created", contratAssurance)
}


//GetContratAssuranceByID method to get an asset by id
func (t *AcheteurAssurance) GetAContratAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetContratAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "asset")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "asset")
	}
	assetAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(assetAsBytes)
}
