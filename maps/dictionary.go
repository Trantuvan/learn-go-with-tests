package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorNotFound  = "could not find the word you were looking for"
	ErrorWordExist = "cannot add word because it already exists"
)

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}
	return "", errors.New(ErrorNotFound)
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	if err != nil {
		d[key] = val
		return nil
	}
	return errors.New(ErrorWordExist)
}
