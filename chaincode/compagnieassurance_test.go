package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewCompagnieAssurance(t *testing.T, stub *shim.MockStub, code string,
	nom string, contact string, adresse string) {
	displayNewTest("Create CompagnieAssurance Test When CompagnieAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateCompagnieAssurance"),
		[]byte(code), []byte(nom), []byte(contact), []byte(adresse)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingCompagnieAssurance(t *testing.T, stub *shim.MockStub, code string) {
	displayNewTest("Get Existing CompagnieAssurance " + code + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetCompagnieAssuranceByID"), []byte(code)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := CompagnieAssurance{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.Code != code {
		t.Fail()
	}
}

func TestCreateCompagnieAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	/* checkCreateNewOrganization(t, stub, "O0", "SCHAIN", "Scorechain", "ISP") */
}

func TestGetCompagnieAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingCompagnieAssurance(t, stub, "O0")
}
