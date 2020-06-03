package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string,
	nom string, contact string, adresse string, passportid string, visaid string) {
	displayNewTest("Create AcheteurAssurance Test When AcheteurAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateAcheteurAssurance"),
		[]byte(uuid), []byte(nom), []byte(contact), []byte(adresse), []byte(passportid), []byte(visaid)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing AcheteurAssurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetAcheteurAssuranceByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := AcheteurAssurance{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.UUID != uuid {
		t.Fail()
	}
}

func TestCreateAcheteurAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkGetExistingAcheteurAssurance(t, stub, "O0")
	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
}

func TestGetAcheteurAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkGetExistingAcheteurAssurance(t, stub, "O0")
}
