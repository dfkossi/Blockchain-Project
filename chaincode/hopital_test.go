package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewHopital(t *testing.T, stub *shim.MockStub, code string,
	nom string, contact string, adresse string) {
	displayNewTest("Create Hopital Test When Hopital does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateHopital"),
		[]byte(code), []byte(nom), []byte(contact), []byte(adresse)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingHopital(t *testing.T, stub *shim.MockStub, code string) {
	displayNewTest("Get Existing Hopital " + code + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetHopitalByID"), []byte(code)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := Hopital{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.Code != code {
		t.Fail()
	}
}

func TestCreateHopital(h *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewHopital(h, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	/* checkCreateNewOrganization(t, stub, "O0", "SCHAIN", "Scorechain", "ISP") */
}

func TestGetHopitalByKey(h *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewHopital(h, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingHopital(h, stub, "O0")
}
