package maddr

import (
	"fmt"
	"strings"

	"github.com/TBD54566975/dap-go/maddr/urn"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

const (
	MoneyAddrKind = "MoneyAddress"
)

type MoneyAddress struct {
	ID       string
	URN      urn.URN
	Currency string
	Protocol string
	PSS      string
}

func FromDIDService(svc didcore.Service) ([]MoneyAddress, error) {
	if svc.Type != MoneyAddrKind {
		return nil, fmt.Errorf("invalid service type: %s", svc.Type)
	}

	maddrs := make([]MoneyAddress, len(svc.ServiceEndpoint))
	for i, se := range svc.ServiceEndpoint {
		maddr, err := Parse(svc.ID, se)
		if err != nil {
			return nil, fmt.Errorf("invalid money address: %w", err)
		}
		maddrs[i] = *maddr
	}
	return maddrs, nil
}

func Parse(serviceID string, maddrURN string) (*MoneyAddress, error) {
	urn, err := urn.Parse(maddrURN)
	if err != nil {
		return nil, fmt.Errorf("invalid money address: %w", err)
	}

	delimIDX := strings.IndexRune(urn.NSS, ':')
	if delimIDX == -1 {
		return nil, fmt.Errorf("invalid money address. expected urn:[currency]:[protocol]:[pss]. got %s", maddrURN)
	}

	return &MoneyAddress{
		URN:      urn,
		ID:       serviceID,
		Currency: urn.NID,
		Protocol: urn.NSS[:delimIDX],
		PSS:      urn.NSS[delimIDX+1:],
	}, nil
}

func MustParse(serviceID string, maddrURN string) MoneyAddress {
	maddr, err := Parse(serviceID, maddrURN)
	if err != nil {
		panic(`maddr: Parse(` + maddrURN + `): ` + err.Error())
	}
	return *maddr
}
