package usecase

import (
	"go-simple/config"
	inNats "go-simple/config/nats"
	"go-simple/model"
	"go-simple/module/v1/article/repo"
)

// get list article
func ArticleList(conf config.Configuration) (users []model.Article, err error) {
	db := conf.MysqlDB
	return repo.GetArticleList(db)
}

// get detail article
func ArticleDetail(conf config.Configuration, artId int) (article model.Article, err error) {
	db := conf.MysqlDB
	return repo.GetArticleDetail(db, artId)
}

// create new article
func ArticleNew(conf config.Configuration, article *model.Article) (user model.Article, err error) {
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	sqlResult, err := repo.CreateNewArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	id, _ := sqlResult.LastInsertId()
	article.Id = int(id)

	tx.Commit()

	// send to event
	if err = inNats.Publish("create.article", &article); err != nil {
		return
	}

	return *article, nil
}

// update article
func ArticleUpdate(conf config.Configuration, article *model.Article) (user model.Article, err error) {

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.UpdateArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *article, nil
}

// delete article
func ArticleDelete(conf config.Configuration, article *model.Article) (user model.Article, err error) {
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.DeleteArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *article, nil
}
