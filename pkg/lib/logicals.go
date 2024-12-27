package lib

// Coalesce will check if the first value is an empty string and return v2 if it is empty, otherwise it will return v1.
func Coalesce(v1, v2 string) string {
	if v1 != "" {
		return v1
	}
	return v2
}

// IIF will return v1 if the condition is true, otherwise it will return v2.
func IIF[T any](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}
	return v2
}
