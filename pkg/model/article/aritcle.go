package article

import (
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	ID    int
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
