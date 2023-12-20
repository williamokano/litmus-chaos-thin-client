package utils

func StringCoalesce(values ...string) string {
	for _, val := range values {
		if val != "" {
			return val
		}
	}

	return ""
}
