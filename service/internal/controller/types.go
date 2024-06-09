package controller

import (
	"fmt"
	"strconv"
	"strings"
)

// CsvInt is a custom type that can be used to unmarshal a comma-separated list of integers from a query parameter
type CsvInt struct {
	value []int
}

func (idl *CsvInt) UnmarshalParam(param string) error {
	parts := strings.Split(param, ",")
	for _, part := range parts {
		intPart, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("\"%s\" is excepted to be a comma-separated list of integers; invalid integer \"%s\"", param, part)
		}
		idl.value = append(idl.value, intPart)
	}
	return nil
}
