package repo

import (
	"database/sql"
	"go-simple/config/database"
	"go-simple/model"

	"github.com/google/martian/log"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

// get user list
func GetPaymentList(sqlx *sqlx.DB) (pay []model.Payment, err error) {
	pay = make([]model.Payment, 0)

	// sql builder
	st := sqlbuilder.NewStruct(new(model.Payment))
	sb := st.SelectFrom(model.TablePayment)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return pay, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return pay, err
		}
		return pay, err
	}

	for rows.Next() {
		var py model.Payment
		if err := rows.Scan(st.Addr(&py)...); err != nil {
			log.Errorf("error : %v", err)
			continue
		}

		pay = append(pay, py)
	}

	return
}

// get user detail
func GetUserDetail(sqlx *sqlx.DB, userId int) (user model.User, err error) {
	var ModelUser model.User

	// sql builder
	st := sqlbuilder.NewStruct(ModelUser)
	sb := st.SelectFrom(model.TabelUser)
	sb.Where(
		sb.Equal("id", userId),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return user, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return user, err
		}
		return user, err
	}

	row.Scan(st.Addr(&user)...)

	return
}

// create payment
func CreatePayment(tx *sql.Tx, p *model.Payment) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Payment{})
	sb := st.InsertIntoForTag(model.TablePayment, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}
