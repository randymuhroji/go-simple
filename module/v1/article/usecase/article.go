package usecase

import (
	"kumparan/config"
	inNats "kumparan/config/nats"
	"kumparan/model"
	"kumparan/module/v1/article/repo"
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
