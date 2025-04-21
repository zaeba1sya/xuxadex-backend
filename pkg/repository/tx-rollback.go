package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

func RollbackTx(tx *sql.Tx, mainErr error) error {
	if err := tx.Rollback(); err != nil {
		return errors.New(fmt.Sprintf("%s. Tx can't rollback.", mainErr.Error()))
	}

	return mainErr
}
