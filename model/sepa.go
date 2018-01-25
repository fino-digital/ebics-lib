package model

type SEPA struct {
	MsgID                   string    `json:"msgid"`
	InitgPty                string    `json:"initParty"`
	PaymentInfoID           string    `json:"paymentInfoid"`
	Creditor                Creditor  `json:"creditor"`
	Debitors                []Debitor `json:"debitor"`
	RequestedCollectionDate string    `json:"RequestedDate,omitempty"`
}

type Creditor struct {
	Name          string         `json:"name"`
	PrivateID     string         `json:"privateid"`
	PostalAddress *PostalAddress `json:"postal,omitempty"`
	IBAN          string         `json:"iban"`
	BIC           string         `json:"bic,omitempty"`
	Currency      string         `json:"currency"`
}

type Debitor struct {
	Name           string  `json:"name"`
	InstructionsID string  `json:"instructionid,omitempty"`
	EndToEndID     string  `json:"endToEndID"`
	MandateID      string  `json:"mandateid"`
	IBAN           string  `json:"iban"`
	BIC            string  `json:"bic,omitempty"`
	Currency       string  `json:"currency"`
	Amount         float64 `json:"amount"`
	Purpose        string  `json:"purpose"`
	DateOfSign     string  `json:"dateOfSign"`
}

type PostalAddress struct {
	Country     string   `json:"country"`
	AddressLine []string `json:"addressLine"`
}
