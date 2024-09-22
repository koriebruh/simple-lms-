package helper

import (
	"fmt"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB, err *error) {
	// Tangkap panic jika ada
	defer func() {
		if r := recover(); r != nil {
			// Jika terjadi panic, rollback transaksi
			tx.Rollback()
			*err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// Jika tidak ada error, commit transaksi, jika ada rollback
	if *err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
