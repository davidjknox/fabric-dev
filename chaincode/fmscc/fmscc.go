package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// Asset definition - The ledger will store an array of call detail records (CDR)

type CDR struct {
	Timestamp string `json:"timestamp"`
	Calling   string `json:"calling"`
	Called    string `json:"called"`
	Duration  string `json:"duration"`
	Value     string `json:"value"`
}

// Main

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

// Init - Initialise the chaincode

func (fms *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("FMS Starting Up...")

	var cdrs []CDR
	bytes, err := json.Marshal(cdrs)

	if err != nil {
		return shim.Error("Error initialising FMS!")
	}

	err = stub.PutState("FMS", bytes)

	return shim.Success(nil)
}

// Invoke - Invocations entry point

func (fms *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("Invoking... " + function)
	if function == "Init" {
		return fms.Init(stub)
	} else if function == "AddCdr" {
		return fms.AddCdr(stub, args)
	}

	return shim.Error("A function named " + function + " doesn't exist!")
}

// AddCdr - Used to add a new CDR

func (fms *SimpleChaincode) AddCdr(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	bytes, err := stub.GetState("FMS")

	if err != nil {
		return shim.Error("Unable to get CDRs")
	}

	var cdr CDR

	// Build JSON values
	timestamp := "\"timestamp\": \"" + args[0] + "\", "
	calling := "\"calling\": \"" + args[1] + "\", "
	called := "\"called\": \"" + args[2] + "\", "
	duration := "\"duration\": \"" + args[3] + "\", "
	value := "\"value\": \"" + args[4] + "\""

	// Make into a complete JSON string
	// Decode into a single CDR item
	content := "{" + timestamp + calling + called + duration + value + "}"
	err = json.Unmarshal([]byte( content ), &cdr)
	fmt.Printf("Query response content: %s\n", content)
	var cdrs []CDR

	// Decode JSON into CDR array
	// Append the new CDR
	err = json.Unmarshal(bytes, &cdrs)
	cdrs = append(cdrs, cdr)

	var threshold = 50.0

	// Check for value greater than £50 and alert if true
	if val, err := strconv.ParseFloat(args[3], 64); err == nil {
		if (val) > threshold {
			fmt.Println("ALERT: Value exceeds threshold!")
		}
	}

	// Encode as JSON
	// Put back on the block
	bytes, err = json.Marshal(cdrs)
	err = stub.PutState("FMS", bytes)
	return shim.Success(nil)
}
