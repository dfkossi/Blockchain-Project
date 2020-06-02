package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewAcheteurAssurance(a *testing.T, stub *shim.MockStub, code string,
	nom string, contact string, adresse string, passportid string, visaid string) {
	displayNewTest("Create AcheteurAssurance Test When AcheteurAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateAcheteurAssurance"),
		[]byte(code), []byte(nom), []byte(contact), []byte(adresse), []byte(passportid), []byte(visaid)})

	if response.Status != shim.OK || response.Payload == nil {
		a.Fail()
	}
}

func checkGetExistingAcheteurAssurance(a *testing.T, stub *shim.MockStub, code string) {
	displayNewTest("Get Existing AcheteurAssurance " + code + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetAcheteurAssuranceByID"), []byte(code)})
	if response.Status != shim.OK || response.Payload == nil {
		a.Fail()
	}

	org := AcheteurAssurance{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.Code != code {
		a.Fail()
	}
}

func TestCreateAcheteurAssurance(a *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(a, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	/* checkCreateNewOrganization(a, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23") */
}

func TestGetAcheteurAssuranceByKey(a *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(a, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkGetExistingAcheteurAssurance(a, stub, "O0")
}
