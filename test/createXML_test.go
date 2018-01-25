package test

import (
	"testing"

	"github.com/fino-digital/ebics-lib/ebics"
	"github.com/fino-digital/ebics-lib/model"
)

func TestCreateXMLFile(t *testing.T) {
	sepa := model.SEPA{
		MsgID:                   "MCSPAD171207X6TZ7LC47Y83",
		InitgPty:                "R+V ALLGEMEINE VERSICHERUNG AG",
		PaymentInfoID:           "1207838J89H24LYT",
		RequestedCollectionDate: "2017-12-11",
		Creditor: model.Creditor{
			Name:      "R+V ALLGEMEINE VERSICHERUNG AG",
			PrivateID: "DE6300100000136090",
			IBAN:      "DE93500604000008512247",
			BIC:       "GENODEFFXXX",
			Currency:  "EUR",
			PostalAddress: &model.PostalAddress{
				Country:     "DE",
				AddressLine: []string{"RAIFFEISENPLATZ 1", "WIESBADEN"},
			},
		},
		Debitors: []model.Debitor{
			model.Debitor{
				Name:           "Kartensparren TEST",
				InstructionsID: "MCSPAD1712072MIMDFSB1890",
				EndToEndID:     "E2ESPAD171207YVS51EW8RSDD",
				MandateID:      "Karten R+V",
				IBAN:           "DE33500604000000012247",
				BIC:            "GENODEFFXXX",
				Currency:       "EUR",
				Amount:         1.00,
				Purpose:        "Test Kartensparen",
				DateOfSign:     "2017-12-07",
			}},
	}

	bytes, err := ebics.BuildSEPAXMLFile(sepa)
	if err != nil {
		t.Fatal(err)
	}

	println(string(bytes))
}
