package maddr

// MoneyAddressFilter accepts a maddr.MoneyAddress, tests if it matches the filter condition, and returns either a more
// specific money address type, or nil
type MoneyAddressFilter[T any] func(maddr MoneyAddress) (*T, error)

// Filter applies the MoneyAddressFilter to each element in maddrs, and returns a slice of the MoneyAddresses that
// matched
func Filter[T any](maddrs []MoneyAddress, filter MoneyAddressFilter[T]) ([]T, error) {
	res := make([]T, 0)
	for _, maddr := range maddrs {
		typedMaddr, err := filter(maddr)
		if err != nil {
			return nil, err
		}
		if typedMaddr != nil {
			res = append(res, *typedMaddr)
		}
	}
	return res, nil
}
