package usecase

import (
	"context"
	"encoding/json"
	"kumparan/config"
	"kumparan/model"

	"kumparan/utl/common"

	"github.com/olivere/elastic"
)

// get list article
func ElasticArticleList(conf config.Configuration, request *model.QueryReq, pagination *model.Pagination) (articles []model.Article, err error) {
	var (
		queryBuilder []elastic.Query
		ctx          = context.Background()
		client       = conf.ElasticeConn
	)

	// Filter Author
	if request.Filter.Author != "" {
		filterAuthor := common.DecodeUrlArray(common.SpaceStringsBuilder(request.Filter.Author))
		queryBuilder = append(queryBuilder, elastic.NewTermsQuery("author.keyword", common.SliceStringToInterface(filterAuthor)...))
	}

	// Filter Title
	if request.Filter.Title != "" {
		filterTitle := common.DecodeUrlArray(common.SpaceStringsBuilder(request.Filter.Title))
		queryBuilder = append(queryBuilder, elastic.NewTermsQuery("title.keyword", common.SliceStringToInterface(filterTitle)...))
	}

	articles = make([]model.Article, 0)
	query := elastic.NewBoolQuery().Filter(elastic.NewBoolQuery().Must(queryBuilder...))
	offset := pagination.Offset()

	search := client.Search().Query(query).Index("article").From(offset).Size(pagination.Limit)

	// sort
	if request.Sort.Value == "ASC" || request.Sort.Value == "DESC" {
		switch request.Sort.Key {
		case "author":
			if request.Sort.Value == "ASC" {
				search.SortBy(elastic.NewFieldSort("author.keyword").Asc().SortMode("max"))
			} else {
				search.SortBy(elastic.NewFieldSort("author.keyword").Desc().SortMode("max"))
			}
		case "title":
			if request.Sort.Value == "ASC" {
				search.SortBy(elastic.NewFieldSort("title.keyword").Asc().SortMode("max"))
			} else {
				search.SortBy(elastic.NewFieldSort("title.keyword").Desc().SortMode("max"))
			}
		case "created_at":
			if request.Sort.Value == "ASC" {
				search.SortBy(elastic.NewFieldSort("created_at").Asc().SortMode("max"))
			} else {
				search.SortBy(elastic.NewFieldSort("created_at").Desc().SortMode("max"))
			}

		}
	}

	result, err := search.Do(ctx)
	if err != nil {
		return articles, err
	}

	for _, hit := range result.Hits.Hits {
		var article model.Article

		json.Unmarshal(hit.Source, &article)

		articles = append(articles, article)
	}

	return
}
