package repo

import (
	"database/sql"
	"go-simple/config/database"
	"go-simple/model"

	"github.com/google/martian/log"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

// get order list
func GetOrderList(sqlx *sqlx.DB) (orders []model.OrderView, err error) {
	orders = make([]model.OrderView, 0)

	// sql builder
	st := sqlbuilder.NewStruct(new(model.OrderView))
	sb := st.SelectFrom(model.TableOrder)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return orders, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return orders, err
		}
		return orders, err
	}

	for rows.Next() {
		var odr model.OrderView
		if err := rows.Scan(st.Addr(&odr)...); err != nil {
			log.Errorf("error : %v", err)
			continue
		}

		orders = append(orders, odr)
	}

	return
}

// create product
func CreateOrder(tx *sql.Tx, p *model.Order) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(new(model.Order))
	sb := st.InsertIntoForTag(model.TableOrder, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// get detail order
func GetDetailOrder(sqlx *sqlx.DB, id int) (order model.Order, err error) {

	// sql builder
	st := sqlbuilder.NewStruct(model.Order{})
	sb := st.SelectFrom(model.TableOrder)
	sb.Where(
		sb.Equal("id", id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return order, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return order, err
		}
		return order, err
	}

	row.Scan(st.Addr(&order)...)

	return
}

// update order
func UpdateOrder(tx *sql.Tx, p *model.Order) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Product{})
	sb := st.UpdateForTag(model.TableProduct, "update", *p)

	sb.Where(
		sb.Equal("id", p.Id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}
