package main 

import {
	"encoding/json"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
}

type FicheSoins struct {
	IdFiche string
	IdContrat string
	IdCompagnieAssurance string
	IdHopital string
	CodeAcheteurAssurance string
	DateDebut datetime
	DateFin datetime
	FichierSoins string
	SignatureAcheteur string
	SignatureCompagnie string
	signatureHopital string
}

func makeFicheSoinsFromBytes(stub.ChaincodeStubInterface, bytes []byte) FicheSoins {
	asset := FicheSoins()
	err := json.Unmarshal(bytes, &ficheSoins)
	panicErr(err)
	return ficheSoinsS
}

func makeBytesFromAsset(stub shim.ChaincodeStubInterface, ficheSoins FicheSoins) []byte {
	bytes, err := json.Marshal(ficheSoins)
	panicErr(err)
	return bytes
}

func createAssetOnLedger(stub shim.ChaincodeStubInterface, objectType string, idFiche string, idContrat string, 
	idCompagnieAssurance string, idHopital string, codeAcheteurAssurance string, dateDebut datetime, dateFin datetime,
	fichierSoins string, signatureAcheteur string, signatureCompagnie string, signatureHopital string) {
	
	ficheSoins := FicheSoins(objectType, idFiche, idContrat, idCompagnieA, idHopital, codeAcheteurAssurance,
	dateDebut, dateFin, fichierSoins, signatureAcheteur, signatureCompagnie, signatureHopital)

	ficheSoinsAsJSONBytes := makeBytesFromAsset(stub, ficheSoins)

	uuidIdexKeyFicheSoins := createIndexKey(stub, idFiche, "asset")
	putEntityInLedger(stub, uuidIdexKeyFicheSoins, ficheSoinsAsJSONBytes)
}

func (f *FicheSoins) CreateFicheSoins(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	idFiche := args[0]
	idCompagnieAssurance := args[0]
	idContrat := args[0]
	idHopital := args[0]
	codeAcheteurAssurance := args[0]
	dateDebut := args[0]
	dateFin := args[0]
	ficheSoins := args[0]
	signatureAcheteur := args[0]
	signatureCompagnie := args[0]
	signatureHopital := args[0] 

	uuidIndexKeyFicheSoins := createIndexKey(stub, idFiche, "FicheSoins")
	ficheSoins := CreateFicheSoinsOnLedger(stub, "FicheSoins",
		uuidIndexKeyFicheSoins, idCompagnieAssurance, idContrat, idHopital, codeAcheteurAssurance, 
		dateDebut, dateFin, fichierSoins, signatureAcheteur, signatureCompagnie, signatureHopital)

	return succeed(stub, "Fiche Soins Created", ficheSoins)
}


//GetFicheSoinsByID method to get an asset by id
func (f *FicheSoins) GetFicheSoinsByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetFicheSoinsByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "asset")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "asset")
	}
	assetAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(assetAsBytes)
}

