package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewContratAssurance(c *testing.T, stub *shim.MockStub, idContrat string, idCompagnieAssurance string,
	codeAcheteurAssurance string, dateDebut datetime, dateFin datetime, fichierContrat string, signatureAcheteur string, signatureCompagnie string) {

	displayNewTest("Create ContratAssurance Test When ContratAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateContratAssurance"),
		[]byte(idContrat), []byte(idCompagnieAssurance), []byte(codeAcheteurAssurance), []byte(dateDebut), []byte(dateFin), []byte(fichierContrat), []byte(signatureAcheteur), []byte(signatureCompagnie)})

	if response.Status != shim.OK || response.Payload == nil {
		c.Fail()
	}
}

func checkGetExistingContratAssurance(c *testing.T, stub *shim.MockStub, idContrat string) {
	displayNewTest("Get Existing ContratAssurance " + idContrat + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetContratAssuranceByID"), []byte(idContrat)})
	if response.Status != shim.OK || response.Payload == nil {
		c.Fail()
	}

	ass := ContratAssurance{}
	_ = json.Unmarshal(response.Payload, &ass)

	if ass.IdContrat != idContrat {
		c.Fail()
	}
}

func TestCreateContratAssurance(c *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(c, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	/* checkCreateNewOrganization(c, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00") */
}

func TestGetContratAssuranceByKey(a *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(a, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	checkGetExistingContratAssurance(a, stub, "O0")
}
