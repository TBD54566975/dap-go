package maddr_test

import (
	"testing"

	"github.com/TBD54566975/dap-go/maddr"
	"github.com/alecthomas/assert"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

func TestDecode(t *testing.T) {
	vectors := []struct {
		input            string
		expectedCurrency string
		expectedCSS      string
		err              bool
	}{
		{
			input:            "urn:usdc:eth:0x2345y7432",
			expectedCurrency: "usdc",
			expectedCSS:      "eth:0x2345y7432",
		},
		{
			input:            "urn:btc:addr:m12345677axcv2345",
			expectedCurrency: "btc",
			expectedCSS:      "addr:m12345677axcv2345",
		},
		{
			input:            "urn:btc:lnurl:https://someurl.com",
			expectedCurrency: "btc",
			expectedCSS:      "lnurl:https://someurl.com",
		},
		{
			input:            "urn:btc:spaddr:sp1234abcd5678",
			expectedCurrency: "btc",
			expectedCSS:      "spaddr:sp1234abcd5678",
		},
	}

	for _, v := range vectors {
		t.Run(v.input, func(t *testing.T) {
			did := didcore.Service{
				Type:            maddr.MoneyAddressKind,
				ID:              "didpay",
				ServiceEndpoint: []string{v.input},
			}
			actual, err := maddr.FromDIDService(did)

			if v.err {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, actual)

				assert.Len(t, actual, 1)
				actualMaddr := actual[0]

				assert.Equal(t, v.expectedCurrency, actualMaddr.Currency)
				assert.Equal(t, v.expectedCSS, actualMaddr.CSS)
				assert.Equal(t, v.expectedCurrency, actualMaddr.URN.NID)
				assert.Equal(t, v.expectedCSS, actualMaddr.URN.NSS)
			}
		})
	}
}
