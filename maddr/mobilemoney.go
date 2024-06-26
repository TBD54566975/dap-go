package maddr

import (
	"fmt"
	"strings"
)

// KESMobileMoneyAddress represents a Mobile Money address for the KES currency. e.g. urn:kes:momo:mpesa:254712345678
type KESMobileMoneyAddress struct {
	MoneyAddress
	Carrier string
	Phone   string
}

func NewKESAddress(maddr MoneyAddress) (*KESMobileMoneyAddress, error) {
	if !isKesAndMoMo(maddr) {
		return nil, fmt.Errorf("invalid currency: %s", maddr.Currency)
	}

	delimIDX := strings.IndexRune(maddr.PSS, ':')
	if delimIDX == -1 {
		return nil, fmt.Errorf("invalid momo money address. expected urn:kes:momo:[carrier]:[phone]. got %s", maddr.URN)
	}
	return &KESMobileMoneyAddress{
		MoneyAddress: maddr,
		Carrier:      maddr.PSS[:delimIDX],
		Phone:        maddr.PSS[delimIDX+1:],
	}, nil
}

func KESMoneyAddressFilter(maddr MoneyAddress) (*KESMobileMoneyAddress, error) {
	if isKesAndMoMo(maddr) {
		return NewKESAddress(maddr)
	}
	return nil, nil
}

func isKesAndMoMo(maddr MoneyAddress) bool {
	return maddr.Currency == "kes" && maddr.Protocol == "momo"
}
