package dto

import "database/sql"

func int16FromModel(n sql.NullInt16) *int16 {
	if n.Valid {
		i := n.Int16

		return &i
	}

	return nil
}
