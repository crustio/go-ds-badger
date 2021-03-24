package badger

import "bytes"

func unseal(value []byte) ([]byte, error) {
	if bytes.HasPrefix(value, []byte("crust/")) {
		return value[6:], nil
	}

	return value, nil
}
