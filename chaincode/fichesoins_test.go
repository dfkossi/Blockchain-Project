package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewFicheSoins(f *testing.T, stub *shim.MockStub, idFiche string, idContrat string,
	idCompagnieAssurance string, idHopital string, codeAcheteurAssurance string, dateDebut string, dateFin string,
	fichierSoins string, signatureAcheteur string, signatureCompagnie string, signatureHopital string) {
	displayNewTest("Create FicheSoins Test When FicheSoins does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateFicheSoins"),
		[]byte(idFiche), []byte(idContrat), []byte(idCompagnieAssurance), []byte(idHopital), []byte(codeAcheteurAssurance), []byte(dateDebut),
		[]byte(dateFin), []byte(fichierSoins), []byte(signatureAcheteur), []byte(signatureCompagnie), []byte(signatureHopital)})

	if response.Status != shim.OK || response.Payload == nil {
		f.Fail()
	}
}

func checkGetExistingFicheSoins(f *testing.T, stub *shim.MockStub, idFiche string) {
	displayNewTest("Get Existing FicheSoins " + idFiche + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetFicheSoinsByID"), []byte(idFiche)})
	if response.Status != shim.OK || response.Payload == nil {
		f.Fail()
	}

	ass := FicheSoins{}
	_ = json.Unmarshal(response.Payload, &ass)

	if ass.IdFiche != idFiche {
		f.Fail()
	}
}

func TestCreateFicheSoins(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "SCHAIN", "Scorechain", "ISP", "SCHAIN",
		"SCHAIN", "Scorechain", "ISP", "Scorechain", "ISP", "UU")
	/* checkCreateNewOrganization(t, stub, "O0", "SCHAIN", "Scorechain", "ISP") */
}

func TestGetFicheSoinsByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "SCHAIN", "Scorechain", "ISP", "SCHAIN",
		"SCHAIN", "Scorechain", "ISP", "Scorechain", "ISP", "UU")
	checkGetExistingFicheSoins(t, stub, "O0")
}
