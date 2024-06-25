package maddr

import (
	"fmt"
	"strings"

	"github.com/TBD54566975/dap-go/maddr/urn"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

const (
	MoneyAddrKind = "maddr"
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
		urn, err := urn.Parse(se)
		if err != nil {
			return nil, fmt.Errorf("invalid money address: %w", err)
		}

		delimIDX := strings.IndexRune(urn.NSS, ':')
		if delimIDX == -1 {
			return nil, fmt.Errorf("invalid money address. expected urn:[currency]:[protocol]:[pss]. got %s", se)
		}

		maddr := MoneyAddress{
			URN:      urn,
			ID:       svc.ID,
			Currency: urn.NID,
			Protocol: urn.NSS[:delimIDX],
			PSS:      urn.NSS[delimIDX+1:],
		}

		maddrs[i] = maddr
	}

	return maddrs, nil
}
