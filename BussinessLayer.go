package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func CreateDatabase(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	//Create table "ContractDetails"
	err = stub.CreateTable("ContractDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ContractId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "OrderId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentCommitment", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "PaymentConfirmation", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "InformationCounterparty", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "ForfeitingInvoice", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "ContractCreateDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentDueDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "InvoiceStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ContractStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "DeliveryStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ContractDetails table.")
	}

	/*
		//Create table "OrderDetails"
		err = stub.CreateTable("OrderDetails", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "OrderId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "ArticleList", Type: shim.ColumnDefinition_BYTES, Key: false},
			&shim.ColumnDefinition{Name: "BuyerId", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "SellerId", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "ShipmentId", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "TotalAmount", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating OrderDetails table.")
		}

		//Create table "SellerDetails"
		err = stub.CreateTable("SellerDetails", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "SellerId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "SellerName", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "SellerBank", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "SellerAddress", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating SellerDetails table.")
		}

		//Create table "BuyerDetails"
		err = stub.CreateTable("BuyerDetails", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "BuyerId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "BuyerName", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "BuyerBank", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "BuyerAddress", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating BuyerDetails table.")
		}

		//Create table "ShipmentDetails"
		err = stub.CreateTable("ShipmentDetails", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "ShipmentId", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "ShipmentCompany", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "ShipmentStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating ShipmentDetails table.")
		}*/
	return nil, nil
}

func InsertContractDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var ContractDetails Contract
	/*var OrderDetails Order
	var SellerDetails Seller
	var BuyerDetails Buyer
	var ShipmentDetails Shipment*/
	var err error
	var ok bool

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 argument")
	}

	//get data from middle layer
	log.Print("************************** 1) *********************************************")
	json.Unmarshal([]byte(args[0]), &ContractDetails)
	log.Print("************************** 2) *********************************************")
	jsonAsBytes, _ := json.Marshal(ContractDetails)
	log.Print("************************** 3) *********************************************")
	log.Print(jsonAsBytes)
	log.Print("************************** 4) *********************************************")
	log.Print(ContractDetails.ContractId)
	log.Print(ContractDetails.OrderId)
	log.Print("************************** 5) *********************************************")
	/*json.Unmarshal([]byte(args[1]), &OrderDetails)
	json.Unmarshal([]byte(args[2]), &SellerDetails)
	json.Unmarshal([]byte(args[3]), &BuyerDetails)
	json.Unmarshal([]byte(args[4]), &ShipmentDetails)

	ArticlesAsBytes, _ := json.Marshal(OrderDetails.ArticleList)*/
	//save data into blockchain
	ok, err = stub.InsertRow("ContractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractId}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.OrderId}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentCommitment}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentConfirmation}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.InformationCounterparty}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.ForfeitingInvoice}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractCreateDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentDueDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.InvoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.DeliveryStatus}},
		},
	})
	if !ok && err == nil {
		return nil, errors.New("Error in adding ContractDetails record.")
	}

	/*ok, err = stub.InsertRow("OrderDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.OrderId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: ArticlesAsBytes}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.BuyerId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.SellerId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.ShipmentId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.TotalAmount}},
		},
	})
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record.")
	}

	ok, err = stub.InsertRow("SellerDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: SellerDetails.SellerId}},
			&shim.Column{Value: &shim.Column_String_{String_: SellerDetails.SellerName}},
			&shim.Column{Value: &shim.Column_String_{String_: SellerDetails.SellerBank}},
			&shim.Column{Value: &shim.Column_String_{String_: SellerDetails.SellerAddress}},
		},
	})
	if !ok && err == nil {
		return nil, errors.New("Error in adding SellerDetails record.")
	}

	ok, err = stub.InsertRow("BuyerDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: BuyerDetails.BuyerId}},
			&shim.Column{Value: &shim.Column_String_{String_: BuyerDetails.BuyerName}},
			&shim.Column{Value: &shim.Column_String_{String_: BuyerDetails.BuyerBank}},
			&shim.Column{Value: &shim.Column_String_{String_: BuyerDetails.BuyerAddress}},
		},
	})
	if !ok && err == nil {
		return nil, errors.New("Error in adding BuyerDetails record.")
	}

	ok, err = stub.InsertRow("ShipmentDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentId}},
			&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentCompany}},
			&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentStatus}},
		},
	})
	if !ok && err == nil {
		return nil, errors.New("Error in adding ShipmentDetails record.")
	}*/

	return nil, nil
}

/*func UpdateContractDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	UserId := args[0]
	UserName := args[1]
	BankName := args[2]
	KycDoc := args[3]
	CurrentDate := time.Now().Local()
	CreateDate := CurrentDate.Format("02-01-2006")
	ValidTillDate := CurrentDate.AddDate(2, 0, 0).Format("02-01-2006")

	ok, err := stub.ReplaceRow("ContractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: UserId}},
			&shim.Column{Value: &shim.Column_String_{String_: UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: BankName}},
			&shim.Column{Value: &shim.Column_String_{String_: CreateDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ValidTillDate}},
			&shim.Column{Value: &shim.Column_String_{String_: KycDoc}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding KYC record.")
	}
	return nil, nil
}

func RegisterMerchant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var MerchantID string    // Entities
	var TxnTokenQty int // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting <MerchantID, Token Name, Token Quantity, Description>")
	}

	// Initialize the chaincode
	MerchantID = args[0]
	TxnTokenName := args[1]
	TxnDescription := args[3]
	TxnTokenQty, err = strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	fmt.Printf("Merchant has = %d\n", TxnTokenQty)

	// Write the state to the ledger
	err = stub.PutState(MerchantID, []byte(strconv.Itoa(TxnTokenQty)))
	if err != nil {
		return nil, err
	}

	txId := stub.GetTxID()
	TxTimestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}

	// create transaction object as payload
	txObj := Transaction{txId, MerchantID, "NA", TxnTokenName, TxnTokenQty, TxnDescription, registerMerchantTxn, TxTimestamp}

	return AddTxToMerchantTxHistory(stub, MerchantID, txObj)
}

func AddTxToMerchantTxHistory(stub shim.ChaincodeStubInterface, MerchantID string, txObj Transaction) ([]byte, error) {

	var merchantTxs []Transaction
	var requiredObj MerchantTransactions
	var objFound bool
	var index int

	merchantIndexTxsAsBytes, err := stub.GetState(merchantIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get merchant Transactions")
	}
	var merchantTxsObjs []MerchantTransactions
	json.Unmarshal(merchantIndexTxsAsBytes, &merchantTxsObjs)

	length := len(merchantTxsObjs)
	// iterate
	for i := 0; i < length; i++ {
		if MerchantID == merchantTxsObjs[i].Name {
			requiredObj = merchantTxsObjs[i]
			objFound = true
			index = i
			break
		}
	}

	if !objFound {
		merchantTxs = append(merchantTxs, txObj)
		merchantTxsObj := MerchantTransactions{MerchantID, merchantTxs}
		merchantTxsObjs = append(merchantTxsObjs, merchantTxsObj)
	} else {
		fmt.Printf("New Transaction object is appended in Merchant : %s 's Transaction History \n", MerchantID)
		merchantTxs = requiredObj.Transactions
		merchantTxs = append(merchantTxs, txObj)
		requiredObj.Transactions = merchantTxs
		merchantTxsObjs[index] = requiredObj

	}

	jsonAsBytes, _ := json.Marshal(merchantTxsObjs)
	err = stub.PutState(merchantIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, errors.New("Failed to update merchant Indexes")
	}

	return nil, nil
}

func RegisterConsumer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var ConsumerID string    // Entities
	var TxnTokenQty int // Asset holdings
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting <ConsumerID, Token Quantity, Description>")
	}

	// Initialize the chaincode
	ConsumerID = args[0]
	TxnDescription := args[2]
	TxnTokenQty, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	fmt.Printf("Consumer has = %d\n", TxnTokenQty)

	// Write the state to the ledger
	err = stub.PutState(ConsumerID, []byte(strconv.Itoa(TxnTokenQty)))
	if err != nil {
		return nil, err
	}

	txId := stub.GetTxID()
	TxTimestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}

	// create transaction object as payload
	txObj := Transaction{txId, "NA", ConsumerID, "NA", TxnTokenQty, TxnDescription, registerConsumerTxn, TxTimestamp}

	return AddTxToConsumerTxHistory(stub, ConsumerID, txObj)
}

func AddTxToConsumerTxHistory(stub shim.ChaincodeStubInterface, ConsumerID string, txObj Transaction) ([]byte, error) {

	var consumerTxs []Transaction
	var requiredObj ConsumerTransactions
	var objFound bool
	var index int

	consumerIndexTxsAsBytes, err := stub.GetState(consumerIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get consumer Transactions")
	}
	var consumerTxsObjs []ConsumerTransactions
	json.Unmarshal(consumerIndexTxsAsBytes, &consumerTxsObjs)

	length := len(consumerTxsObjs)
	// iterate
	for i := 0; i < length; i++ {
		if ConsumerID == consumerTxsObjs[i].Name {
			requiredObj = consumerTxsObjs[i]
			objFound = true
			index = i
			break
		}
	}
	if !objFound {
		consumerTxs = append(consumerTxs, txObj)
		consumerTxsObj := ConsumerTransactions{ConsumerID, consumerTxs}
		consumerTxsObjs = append(consumerTxsObjs, consumerTxsObj)
	} else {
		fmt.Printf("New Transaction object is appended in Consumer : %s 's Transaction History \n", ConsumerID)
		consumerTxs = requiredObj.Transactions
		consumerTxs = append(consumerTxs, txObj)
		requiredObj.Transactions = consumerTxs
		consumerTxsObjs[index] = requiredObj
	}

	jsonAsBytes, _ := json.Marshal(consumerTxsObjs)
	err = stub.PutState(consumerIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, errors.New("Failed to update consumer Indexes")
	}

	return nil, nil
}

func AddTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var MerchantID, ConsumerID string    // Entities
	var MerchantVal, ConsumerVal int // Asset holdings
	var TxnTokenQty int          // Transaction value
	var err error

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting <MerchantID, ConsumerID, Token Name, Token Quantity, Description>")
	}

	MerchantID = args[0]
	ConsumerID = args[1]
	TxnTokenName := args[2]
	TxnDescription := args[4]
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(MerchantID)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	MerchantVal, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(ConsumerID)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	ConsumerVal, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	TxnTokenQty, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Invalid transaction amount, expecting a integer value")
	}
	MerchantVal = MerchantVal - TxnTokenQty
	ConsumerVal = ConsumerVal + TxnTokenQty
	fmt.Printf("MerchantVal = %d, ConsumerVal = %d\n", MerchantVal, ConsumerVal)

	// Write the state back to the ledger
	err = stub.PutState(MerchantID, []byte(strconv.Itoa(MerchantVal)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(ConsumerID, []byte(strconv.Itoa(ConsumerVal)))
	if err != nil {
		return nil, err
	}

	txId := stub.GetTxID()
	TxTimestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}

	// create transaction object as payload
	txObj := Transaction{txId, MerchantID, ConsumerID, TxnTokenName, TxnTokenQty, TxnDescription, addTransaction, TxTimestamp}

	_, err = AddTxToMerchantTxHistory(stub, MerchantID, txObj)
	if err != nil {
		return nil, err
	}

	_, err = AddTxToConsumerTxHistory(stub, ConsumerID, txObj)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetMerchantTxHistory(stub shim.ChaincodeStubInterface, MerchantID string) ([]byte, error) {

	var requiredObj MerchantTransactions

	merchantIndexTxsAsBytes, err := stub.GetState(merchantIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get Merchant Transactions")
	}
	var merchantTxObjects []MerchantTransactions
	json.Unmarshal(merchantIndexTxsAsBytes, &merchantTxObjects)
	length := len(merchantTxObjects)

	// iterate
	for i := 0; i < length; i++ {
		obj := merchantTxObjects[i]
		if MerchantID == obj.Name {
			requiredObj = obj
		}
	}

	fmt.Println("required Obj name: %s and input name: %s", requiredObj.Name, MerchantID)
	res, err := json.Marshal(requiredObj)
	if err != nil {
		return nil, errors.New("Failed to Marshal the required Obj")
	}

	return res, nil
}

func GetConsumerTxHistory(stub shim.ChaincodeStubInterface, ConsumerID string) ([]byte, error) {

	var requiredObj ConsumerTransactions

	consumerIndexTxsAsBytes, err := stub.GetState(consumerIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get consumer Transactions")
	}
	var consumerTxObjects []ConsumerTransactions
	json.Unmarshal(consumerIndexTxsAsBytes, &consumerTxObjects)
	length := len(consumerTxObjects)

	// iterate
	for i := 0; i < length; i++ {
		obj := consumerTxObjects[i]
		if ConsumerID == obj.Name {
			requiredObj = obj
		}
	}

	fmt.Println("required Obj name: %s and input name: %s", requiredObj.Name, ConsumerID)

	res, err := json.Marshal(requiredObj)
	if err != nil {
		return nil, errors.New("Failed to Marshal the required Obj")
	}

	return res, nil
}*/
