package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// CsvIDs is a custom type that can be used to unmarshal a comma-separated list of positive integers from a query parameter.
type CsvIDs struct {
	value []int
}

func (idl *CsvIDs) UnmarshalParam(param string) error {
	parts := strings.Split(param, ",")
	for _, part := range parts {
		intPart, err := strconv.Atoi(part)
		if err != nil || intPart < 1 {
			msg := fmt.Sprintf("\"%s\" is excepted to be a comma-separated list of positive integers; invalid value \"%s\"", param, part)
			if err != nil {
				return errors.Wrapf(err, msg)
			}
			return errors.New(msg)
		}
		idl.value = append(idl.value, intPart)
	}
	return nil
}
