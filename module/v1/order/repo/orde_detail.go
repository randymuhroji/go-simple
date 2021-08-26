package repo

import (
	"database/sql"
	"go-simple/config/database"
	"go-simple/model"

	"github.com/huandu/go-sqlbuilder"
)

// create product
func CreateOrderDetail(tx *sql.Tx, p *model.OrderDetail) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.OrderDetail{})
	sb := st.InsertIntoForTag(model.TableOrderDetail, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}
