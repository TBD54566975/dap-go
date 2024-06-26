package maddr_test

import (
	"github.com/TBD54566975/dap-go/maddr"
	"github.com/alecthomas/assert"
	"github.com/tbd54566975/web5-go/dids/didcore"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []string{"urn:kes:momo:mpesa:254712345678", "urn:usdc:eth:0x2345y7432", "urn:btc:addr:m12345677axcv2345", "urn:btc:lnurl:https://someurl.com", "urn:btc:spaddr:sp1234abcd5678"}
	maddrs := parseMoneyAddresses(t, input)
	expectedOut := maddr.KESMobileMoneyAddress{
		MoneyAddress: maddr.MustParse("id", "urn:kes:momo:mpesa:254712345678"),
		Carrier:      "mpesa",
		Phone:        "254712345678",
	}

	out, err := maddr.Filter(maddrs, maddr.KESMoneyAddressFilter)

	assert.NoError(t, err)
	assert.NotNil(t, out)
	assert.Len(t, out, 1)
	actualMaddr := out[0]
	if !reflect.DeepEqual(actualMaddr, expectedOut) {
		t.Errorf("Filter() got = %v, want %v", expectedOut, actualMaddr)
	}
}

func parseMoneyAddresses(t *testing.T, maddrs []string) []maddr.MoneyAddress {
	did := didcore.Service{
		Type:            maddr.MoneyAddrKind,
		ID:              "id",
		ServiceEndpoint: maddrs,
	}
	service, err := maddr.FromDIDService(did)
	assert.NoError(t, err)
	return service
}
