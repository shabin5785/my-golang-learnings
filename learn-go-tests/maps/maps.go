package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("missing value")

func (d Dictionary) Search(word string) (string, error) {
	v, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) {
	d[k] = v
}
