package maps

import "errors"

type Dictionary map[string]string

var ErrorNotFound = "could not find the word you were looking for"

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}
	return "", errors.New(ErrorNotFound)
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
