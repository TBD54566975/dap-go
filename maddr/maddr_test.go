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
		expectedProtocol string
		expectedPSS      string
		err              bool
	}{
		{
			input:            "urn:usdc:eth:0x2345y7432",
			expectedCurrency: "usdc",
			expectedProtocol: "eth",
			expectedPSS:      "0x2345y7432",
		},
		{
			input:            "urn:btc:addr:m12345677axcv2345",
			expectedCurrency: "btc",
			expectedProtocol: "addr",
			expectedPSS:      "m12345677axcv2345",
		},
		{
			input:            "urn:btc:lnurl:https://someurl.com",
			expectedCurrency: "btc",
			expectedProtocol: "lnurl",
			expectedPSS:      "https://someurl.com",
		},
		{
			input:            "urn:btc:spaddr:sp1234abcd5678",
			expectedCurrency: "btc",
			expectedProtocol: "spaddr",
			expectedPSS:      "sp1234abcd5678",
		},
	}

	for _, v := range vectors {
		t.Run(v.input, func(t *testing.T) {
			did := didcore.Service{
				Type:            maddr.MoneyAddrKind,
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
				assert.Equal(t, v.expectedProtocol, actualMaddr.Protocol)
				assert.Equal(t, v.expectedPSS, actualMaddr.PSS)
				assert.Equal(t, v.expectedCurrency, actualMaddr.URN.NID)
				assert.Equal(t, v.expectedProtocol+":"+v.expectedPSS, actualMaddr.URN.NSS)
			}
		})
	}
}
