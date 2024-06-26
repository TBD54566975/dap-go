package maddr_test

import (
	"github.com/TBD54566975/dap-go/maddr"
	"github.com/alecthomas/assert"
	"testing"
)

func TestNewKESAddress(t *testing.T) {
	tests := []struct {
		input     maddr.MoneyAddress
		expected  *maddr.KESMobileMoneyAddress
		expectErr bool
	}{
		{
			input: maddr.MustParse("id", "urn:kes:momo:mpesa:254712345678"),
			expected: &maddr.KESMobileMoneyAddress{
				MoneyAddress: maddr.MustParse("id", "urn:kes:momo:mpesa:254712345678"),
				Carrier:      "mpesa",
				Phone:        "254712345678",
			},
			expectErr: false,
		},
		{
			input:     maddr.MustParse("id", "urn:kes:momo:mpesa"),
			expectErr: true,
		},
		{
			input:     maddr.MustParse("id", "urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip"),
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input.URN.URN, func(t *testing.T) {
			actual, err := maddr.NewKESAddress(tt.input)

			if (err != nil) != tt.expectErr {
				t.Errorf("NewKESAddress() error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
