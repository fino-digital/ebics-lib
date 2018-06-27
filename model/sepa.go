package model

type SEPA struct {
	MsgID                   string    `json:"msgid"` // Point-to-point reference of the instructing party for the next party in the message chain to uniquely identify the message (file). Can be generated with ebics.RestrictedIdentificationSEPA1()
	InitgPty                string    `json:"initParty"`
	PaymentInfoID           string    `json:"paymentInfoid"` // Reference to uniquely identify the following collector. Can be generated with ebics.RestrictedIdentificationSEPA1()
	BatchBook               bool      `json:"batchbook"`     // Indicator that states whether it is a collective entry (true) or a single entry (false). [0..1]
	Creditor                Creditor  `json:"creditor"`
	Debitors                []Debitor `json:"debitor"`
	RequestedCollectionDate string    `json:"RequestedDate,omitempty"`
}

type Creditor struct {
	Name          string         `json:"name"`
	CreditorID    string         `json:"creditorid"` // Creditor Identifier (CI)
	PostalAddress *PostalAddress `json:"postal,omitempty"`
	IBAN          string         `json:"iban"`
	BIC           string         `json:"bic,omitempty"`
	Currency      string         `json:"currency"`
}

type Debitor struct {
	Name           string  `json:"name"`
	InstructionsID string  `json:"instructionid,omitempty"` // Unique reference of the direct debit recipient to his bank (point-to-point reference). Can be generated with ebics.RestrictedIdentificationSEPA1()
	EndToEndID     string  `json:"endToEndID"`              // This reference is passed unchanged through the entire chain to the payer (end-to-end reference).
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
