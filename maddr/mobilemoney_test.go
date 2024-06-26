package maddr_test

import (
	"github.com/TBD54566975/dap-go/maddr"
	"github.com/alecthomas/assert"
	"reflect"
	"testing"
)

func TestNewKESAddress(t *testing.T) {
	tests := []struct {
		input   *maddr.MoneyAddress
		want    *maddr.KESMobileMoneyAddress
		wantErr bool
	}{
		{
			input: newMoneyAddress(t, "urn:kes:momo:mpesa:254712345678"),
			want: &maddr.KESMobileMoneyAddress{
				MoneyAddress: *newMoneyAddress(t, "urn:kes:momo:mpesa:254712345678"),
				Carrier:      "mpesa",
				Phone:        "254712345678",
			},
			wantErr: false,
		},
		{
			input:   newMoneyAddress(t, "urn:kes:momo:mpesa"),
			wantErr: true,
		},
		{
			input:   newMoneyAddress(t, "urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input.URN.URN, func(t *testing.T) {
			got, err := maddr.NewKESAddress(*tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewKESAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKESAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func newMoneyAddress(t *testing.T, maddrURN string) *maddr.MoneyAddress {
	m, err := maddr.FromURN("didpay", maddrURN)
	assert.NoError(t, err)
	return m
}
