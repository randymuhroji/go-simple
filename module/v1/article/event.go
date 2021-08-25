package article

import (
	"context"
	inElastic "kumparan/config/elastic"
	inNats "kumparan/config/nats"
	"kumparan/model"
	"kumparan/module/v1/article/usecase"

	"github.com/labstack/gommon/log"
	"github.com/nats-io/nats.go"
)

// handle event create article
func (m *Module) EventCreateArticle() {
	var (
		atr model.Article
		ctx = context.Background()
	)

	inNats.Subscription("create.article", func(msg *nats.Msg) {
		if err := inNats.ReadMessage(msg.Data, &atr); err != nil {
			log.Error(err)
		}
		article, err := usecase.ArticleDetail(m.Config, atr.Id)
		if err != nil || article.Id == 0 {
			return
		}
		client := m.Config.ElasticeConn
		if err := inElastic.ValidIndex(client, "article"); err != nil {
			return
		}

		// insert into elastic
		if _, err = client.Index().Index("article").BodyJson(article).Do(ctx); err != nil {
			return
		}
		defer client.CloseIndex("article")
	})

}
