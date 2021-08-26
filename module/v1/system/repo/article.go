package repo

import (
	"database/sql"
	"go-simple/config/database"
	"go-simple/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

// get user list
func GetArticleList(sqlx *sqlx.DB) (articles []model.Article, err error) {
	articles = make([]model.Article, 0)
	var ModelArticle model.Article
	// sql builder
	st := sqlbuilder.NewStruct(ModelArticle)
	sb := st.SelectFrom(model.TableArticle)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return articles, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return articles, err
		}
		return articles, err
	}

	for rows.Next() {
		var usr model.Article
		if err := rows.Scan(st.Addr(&usr)...); err != nil {
			log.Error(err)
			continue
		}

		articles = append(articles, usr)
	}

	return
}

// get detail article by id
func GetArticleDetail(sqlx *sqlx.DB, articleId int) (article model.Article, err error) {
	var ModelArticle model.Article

	// sql builder
	st := sqlbuilder.NewStruct(ModelArticle)
	sb := st.SelectFrom(model.TableArticle)
	sb.Where(
		sb.Equal("id", articleId),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return article, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return article, err
		}
		return article, err
	}

	row.Scan(st.Addr(&article)...)

	return
}

// create new article
func CreateNewArticle(tx *sql.Tx, p *model.Article) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Article{})
	sb := st.InsertIntoForTag(model.TableArticle, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// update article
func UpdateArticle(tx *sql.Tx, p *model.Article) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Article{})
	sb := st.UpdateForTag(model.TableArticle, "update", *p)

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

// delete article
func DeleteArticle(tx *sql.Tx, p *model.Article) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.Article{})
	sb := st.UpdateForTag(model.TableArticle, "delete", *p)
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

// get detail article by param
func GetArticleDetailByParam(sqlx *sqlx.DB, param string, value interface{}) (article model.Article, err error) {
	var ModelArticle model.Article

	// sql builder
	st := sqlbuilder.NewStruct(ModelArticle)
	sb := st.SelectFrom(model.TableArticle)
	sb.Where(
		sb.Equal(param, value),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return article, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return article, err
		}
		return article, err
	}

	row.Scan(st.Addr(&article)...)

	return
}
