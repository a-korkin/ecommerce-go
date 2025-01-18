package models

import (
	"github.com/a-korkin/ecommerce/internal/utils"
	"strconv"
)

type PageParams struct {
	Page  int
	Limit int
}

func NewPageParams(url string) *PageParams {
	pageParams := PageParams{
		Page:  1,
		Limit: 20,
	}
	params := utils.GetQueryParams(url)

	page, ok := params["page"]
	if ok {
		p, err := strconv.Atoi(page)
		if err == nil {
			pageParams.Page = p
		}
	}
	limit, ok := params["limit"]
	if ok {
		l, err := strconv.Atoi(limit)
		if err == nil {
			pageParams.Limit = l
		}
	}

	return &pageParams
}
