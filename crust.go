package badger

import (
	"bytes"

	"github.com/dgraph-io/badger"
)

func crustUnseal(item *badger.Item) ([]byte, error) {
	value, err := item.ValueCopy(nil)
	if err != nil {
		return value, err
	}

	if bytes.HasPrefix(value, []byte("crust/")) {
		return value[6:], nil
	}

	return value, nil
}

func crustGetSize(item *badger.Item) (int, error) {
	value, err := item.ValueCopy(nil)
	if err != nil {
		return 0, err
	}

	if bytes.HasPrefix(value, []byte("crust/")) {
		return len(value) - 6, nil
	}

	return len(value), nil
}
