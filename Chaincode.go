package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Region Chaincode implementation
type DTC_Chaincode struct {
}

func (t *DTC_Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	return CreateDatabase(stub, args)

}

// Add user KYC data in Blockchain
func (t *DTC_Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "InsertContractDetails" {
		// Insert Contract data in blockchain
		return InsertContractDetails(stub, args)
	} /* else if function == "UpdateContractDetails" {
		// Update Contract data in blockchain
		return UpdateKycDetails(stub, args)
	}*/

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *DTC_Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var ContractDetails Contract
	var jsonAsBytes []byte
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting enrollId to query")
	}

	UserId := args[0]

	if UserId == "" {
		table, err := stub.GetTable("ContractDetails")
		if err != nil {
			return nil, errors.New("Failed to query")
		}
		//output := table.String()
		jsonAsBytes, _ = json.Marshal(table)
	} else {
		var columns []shim.Column

		col1 := shim.Column{Value: &shim.Column_String_{String_: UserId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("ContractDetails", columns)
		if err != nil {
			return nil, errors.New("Failed to query")
		}

		ContractDetails.ContractId = row.Columns[0].GetString_()
		ContractDetails.OrderId = row.Columns[1].GetString_()
		ContractDetails.PaymentCommitment = row.Columns[2].GetBool()
		ContractDetails.PaymentConfirmation = row.Columns[3].GetBool()
		ContractDetails.InformationCounterparty = row.Columns[4].GetBool()
		ContractDetails.ForfeitingInvoice = row.Columns[5].GetBool()
		ContractDetails.ContractCreateDate = row.Columns[6].GetString_()
		ContractDetails.PaymentDueDate = row.Columns[7].GetString_()
		ContractDetails.InvoiceStatus = row.Columns[8].GetString_()
		ContractDetails.PaymentStatus = row.Columns[9].GetString_()
		ContractDetails.ContractStatus = row.Columns[10].GetString_()
		ContractDetails.DeliveryStatus = row.Columns[11].GetString_()

		jsonAsBytes, _ = json.Marshal(ContractDetails)
	}

	return jsonAsBytes, nil
}

func main() {
	err := shim.Start(new(DTC_Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
