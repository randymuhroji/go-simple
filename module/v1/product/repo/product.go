package repo

import (
	"database/sql"
	"go-simple/config/database"
	"go-simple/model"

	"github.com/google/martian/log"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

// get product list
func GetProductList(sqlx *sqlx.DB) (products []model.ProductView, err error) {
	products = make([]model.ProductView, 0)

	// sql builder
	st := sqlbuilder.NewStruct(new(model.ProductView))
	sb := st.SelectFrom(model.TableProduct)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return products, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return products, err
		}
		return products, err
	}

	for rows.Next() {
		var prd model.ProductView
		if err := rows.Scan(st.Addr(&prd)...); err != nil {
			log.Errorf("error : %v", err)
			continue
		}

		products = append(products, prd)
	}

	return
}

// get product detail
func GetProductDetail(sqlx *sqlx.DB, id int) (product model.ProductView, err error) {

	// sql builder
	st := sqlbuilder.NewStruct(model.ProductView{})
	sb := st.SelectFrom(model.TableProduct)
	sb.Where(
		sb.Equal("id", id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return product, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("error : %v", err)
			return product, err
		}
		return product, err
	}

	row.Scan(st.Addr(&product)...)

	return
}

// create product
func CreateProduct(tx *sql.Tx, p *model.Product) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Product{})
	sb := st.InsertIntoForTag(model.TableProduct, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// update user
func UpadateProduct(tx *sql.Tx, p *model.Product) (result sql.Result, err error) {
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

// delete product
func DeleteProduct(tx *sql.Tx, p *model.Product) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Product{})
	sb := st.UpdateForTag(model.TableProduct, "delete", *p)
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
