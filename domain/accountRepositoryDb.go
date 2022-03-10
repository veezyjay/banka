package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/veezyjay/banka/errs"
	"github.com/veezyjay/banka/logger"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	newAccountQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.db.Exec(newAccountQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
