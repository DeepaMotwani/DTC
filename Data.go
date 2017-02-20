package main

type Contract struct {
	ContractId              string `json:"contractId"`
	OrderId                 string `json:"orderId"`
	PaymentCommitment       bool   `json:"paymentCommitment"`
	PaymentConfirmation     bool   `json:"paymentConfirmation"`
	InformationCounterparty bool   `json:"informationCounterparty"`
	ForfeitingInvoice       bool   `json:"forfeitingInvoice"`
	ContractCreateDate      string `json:"createDate"`
	PaymentDueDate          string `json:"paymentDueDate"`
	InvoiceStatus           string `json:"paymentStatus"`
	PaymentStatus           string `json:"paymentStatus"`
	ContractStatus          string `json:"contractStatus"`
	DeliveryStatus          string `json:"deliveryStatus"`
}
type Order struct {
	OrderId     string    `json:"orderId"`
	ArticleList []Article `json:"articles"`
	BuyerId     string    `json:"buyerId"`
	SellerId    string    `json:"sellerId"`
	ShipmentId  string    `json:"shipmentId"`
	TotalAmount string    `json:"amount"`
}
type Seller struct {
	SellerId      string `json:"sellerId"`
	SellerName    string `json:"name"`
	SellerBank    string `json:"bank"`
	SellerAddress string `json:"address"`
}
type Buyer struct {
	BuyerId      string `json:"buyerId"`
	BuyerName    string `json:"name"`
	BuyerBank    string `json:"bank"`
	BuyerAddress string `json:"address"`
}
type Shipment struct {
	ShipmentId      string `json:"trackingId"`
	ShipmentCompany string `json:"ShipmentCompany"`
	ShipmentStatus  string `json:"ShipmentStatus"`
}
type Article struct {
	ArticleDescription string `json:"description"`
	ArticleQty         int    `json:"quantity"`
	ArticlePrice       string `json:"amount"`
}

/*type Company struct {
	CompanyName string `json:"name"`
}

type Address struct {
	Line    string `json:"line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}
type Bank struct {
	BankName   string `json:"name"`
	BranchName string `json:"branch"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
}
type Amount struct {
	Currency string `json:"currency"`
	Value    uint64 `json:"value"`
}*/
