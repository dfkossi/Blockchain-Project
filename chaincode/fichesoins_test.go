package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewFicheSoins(t *testing.T, stub *shim.MockStub, uuid string, iDContrat string,
	iDCompagnieAssurance string, iDHopital string, codeAcheteurAssurance string, dateDebut string, dateFin string,
	ficheSoinsPDF string, signatureAcheteur string, signatureCompagnie string, signatureHopital string) {
	displayNewTest("Create FicheSoins Test When FicheSoins does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateFicheSoins"),
		[]byte(uuid), []byte(iDContrat), []byte(iDCompagnieAssurance), []byte(iDHopital), []byte(codeAcheteurAssurance), []byte(dateDebut),
		[]byte(dateFin), []byte(ficheSoinsPDF), []byte(signatureAcheteur), []byte(signatureCompagnie), []byte(signatureHopital)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingFicheSoins(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing FicheSoins " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetFicheSoinsByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	ass := FicheSoins{}
	_ = json.Unmarshal(response.Payload, &ass)

	if ass.UUID != uuid {
		t.Fail()
	}
}

func TestCreateFicheSoins(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "SCHAIN", "Scorechain", "ISP", "SCHAIN",
		"SCHAIN", "Scorechain", "ISP", "Scorechain", "ISP", "UU")
	checkGetExistingFicheSoins(t, stub, "O0")
	checkCreateNewFicheSoins(t, stub, "O0", "SCHAIN", "Scorechain", "ISP", "SCHAIN",
		"SCHAIN", "Scorechain", "ISP", "Scorechain", "ISP", "UU")
}

func TestGetFicheSoinsByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "SCHAIN", "Scorechain", "ISP", "SCHAIN",
		"SCHAIN", "Scorechain", "ISP", "Scorechain", "ISP", "UU")
	checkGetExistingFicheSoins(t, stub, "O0")
}
