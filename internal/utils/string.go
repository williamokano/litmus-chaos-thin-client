package utils

import "fmt"

type StringSlice []string

func (t *StringSlice) String() string {
	return fmt.Sprintf("%v", *t)
}

func (t *StringSlice) Set(value string) error {
	*t = append(*t, value)
	return nil
}

func StringCoalesce(values ...string) string {
	for _, val := range values {
		if val != "" {
			return val
		}
	}

	return ""
}
