package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}
