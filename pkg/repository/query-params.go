package repository

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type sortData struct {
	Field string
	Order string
}

type QueryOpts struct {
	Page   uint32
	Limit  uint32
	Sort   sortData
	Filter string
}

/*
QueryOpts represents basic query parameters for pagination, sorting, and filtering.
Page indicates the current page number (default 1).
Limit defines max number of items per page (default 10).
Sort specifies ordering parameter. (e.g., "id:DESC", "name:ASC")
Filter holds filter conditions. (e.g., "status=active", "price>100")
*/
func ParseQueryOpts(ctx echo.Context) *QueryOpts {
	var page, limit uint32 = 1, 10
	var sortRaw, filter string = ctx.QueryParam("sort"), ctx.QueryParam("filter")

	rawPage, err := strconv.Atoi(ctx.QueryParam("page"))
	if err == nil && rawPage != 0 {
		page = uint32(rawPage)
	}

	rawLimit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err == nil && rawLimit != 0 {
		limit = uint32(rawLimit)
	}

	opts := &QueryOpts{
		Page:  page,
		Limit: limit,
		Sort: sortData{
			Field: "id",
			Order: "ASC",
		},
		Filter: filter,
	}

	if sortRaw != "" {
		sort := strings.Split(sortRaw, ":")
		if len(sort) == 2 {
			opts.Sort = sortData{
				Field: sort[0],
				Order: sort[1],
			}
		}
	}

	return opts
}
