package controller

import (
	"fmt"
	"strconv"
	"strings"
)

// CsvIds is a custom type that can be used to unmarshal a comma-separated list of positive integers from a query parameter
type CsvIds struct {
	value []int
}

func (idl *CsvIds) UnmarshalParam(param string) error {
	parts := strings.Split(param, ",")
	for _, part := range parts {
		intPart, err := strconv.Atoi(part)
		if err != nil || intPart < 1 {
			return fmt.Errorf("\"%s\" is excepted to be a comma-separated list of positive integers; invalid value \"%s\"", param, part)
		}
		idl.value = append(idl.value, intPart)
	}
	return nil
}
