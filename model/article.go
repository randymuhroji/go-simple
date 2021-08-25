package model

const TableArticle = "article"

type Article struct {
	Id        int    `json:"id" db:"id"`
	Author    string `json:"author" db:"article_author" fieldtag:"insert,update"`
	Title     string `json:"title" db:"article_title" fieldtag:"insert,update"`
	Body      string `json:"body" db:"article_body" fieldtag:"insert,update"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}
