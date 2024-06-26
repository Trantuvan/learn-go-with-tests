package maps

func Search(dict map[string]string, key string) string {
	if val, ok := dict[key]; ok {
		return val
	}
	return ""
}
