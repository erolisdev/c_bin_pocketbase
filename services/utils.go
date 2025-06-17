package services

import (
	"fmt"
	"strconv"
)

func PtrValue[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
}

// areFloatsEqual checks if two floats are equal within a given tolerance.
func areFloatsEqual(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < tolerance
}

func parseFloat(sVal *string, fieldName string) (float64, error) {
	if sVal == nil || *sVal == "" {
		return 0.0, fmt.Errorf("%s is missing or empty", fieldName)
	}
	f, err := strconv.ParseFloat(*sVal, 64)
	if err != nil {
		return 0.0, fmt.Errorf("invalid %s value '%s': %w", fieldName, *sVal, err)
	}
	return f, nil
}

func derefString(s *string, defaultValue string) string {
	if s != nil {
		return *s
	}
	return defaultValue
}

func derefInt64(i *int64, defaultValue int64) int64 {
	if i != nil {
		return *i
	}
	return defaultValue
}

func derefFloat64(f *float64, defaultValue float64) float64 {
	if f != nil {
		return *f
	}
	return defaultValue
}

func safeStringFromInt64Ptr(ptr *int64) string {
	if ptr == nil {
		return "" // Veya bir hata döndürün ya da loglayın, duruma göre
	}
	return strconv.FormatInt(*ptr, 10)
}

func ToStringSimple(value any) string {
	if value == nil {
		return "" // veya "nil" tercih edebilirsiniz
	}
	return fmt.Sprint(value)
}

// Helper to convert []string to []interface{} for dbx.In
func convertToInterfaceSlice[T any](slice []T) []interface{} {
	iSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		iSlice[i] = v
	}
	return iSlice
}
