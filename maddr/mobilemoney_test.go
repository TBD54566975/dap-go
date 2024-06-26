package maddr

import (
	"github.com/alecthomas/assert"
	"reflect"
	"testing"
)

func TestNewKESAddress(t *testing.T) {
	tests := []struct {
		maddr   *MoneyAddress
		want    *KESMobileMoneyAddress
		wantErr bool
	}{
		{
			maddr: newMoneyAddress(t, "urn:kes:momo:mpesa:254712345678"),
			want: &KESMobileMoneyAddress{
				MoneyAddress: *newMoneyAddress(t, "urn:kes:momo:mpesa:254712345678"),
				Carrier:      "mpesa",
				Phone:        "254712345678",
			},
			wantErr: false,
		},
		{
			maddr:   newMoneyAddress(t, "urn:kes:momo:mpesa"),
			wantErr: true,
		},
		{
			maddr:   newMoneyAddress(t, "urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.maddr.URN.URN, func(t *testing.T) {
			got, err := NewKESAddress(*tt.maddr)

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

func newMoneyAddress(t *testing.T, maddr string) *MoneyAddress {
	m, err := FromURN("didpay", maddr)
	assert.NoError(t, err)
	return m
}
