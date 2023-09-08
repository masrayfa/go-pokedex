package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		PanicIfError(err)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}