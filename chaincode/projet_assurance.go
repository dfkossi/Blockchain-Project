package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//ProjetAssurance principal chaincode class
type ProjetAssurance struct {
	compagnieAssurance CompagnieAssurance
	hopital            Hopital
}

//Init function to Initiate the chaincode
func (t *ProjetAssurance) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Init")
	return shim.Success([]byte("Init success"))
}

//Invoke function to invoke the chaincode
func (t *ProjetAssurance) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	defer func() {
		if r := recover(); r != nil {
			functionnalError, ok := r.(FunctionnalError)
			if ok {
				shim.Error(functionnalError.errorTag)
			}
			err, ok := r.(error)
			if ok {
				shim.Error(fmt.Sprintf("%v", err))
			} else {
				shim.Error("unknownError")
			}
		}
	}()

	fc, args := stub.GetFunctionAndParameters()

	switch {
	// COMPAGNIEASSURANCE
	case strings.Compare(fc, "CreateCompagnieAssurance") == 0:
		return t.compagnieAssurance.CreateCompagnieAssurance(stub, args)

	case strings.Compare(fc, "GetCompagnieAssuranceByID") == 0:
		return t.compagnieAssurance.GetCompagnieAssuranceByID(stub, args[0])

		// HOPITAL
	case strings.Compare(fc, "CreateHopital") == 0:
		return t.hopital.CreateHopital(stub, args)

	/* case strings.Compare(fc, "GetCompagnieAssuranceByID") == 0:
	return t.compagnieAssurance.GetCompagnieAssuranceByID(stub, args[0]) */

	default:
		return shim.Error("Called function is not defined in the chaincode ")
	}

}

func main() {
	err := shim.Start(new(ProjetAssurance))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
