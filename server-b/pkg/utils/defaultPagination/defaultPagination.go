package defaultpagination

import (
	"strconv"

	"server/pkg/apperror"
	"server/pkg/dto"

	"github.com/gin-gonic/gin"
)

func SetQuery(ctx *gin.Context) (*dto.Query, error) {
	query := ctx.DefaultQuery("q", "")
	sortBy := ctx.DefaultQuery("sortBy", "created_at")
	sort := ctx.DefaultQuery("sort", "asc")
	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		return nil, apperror.NewInvalidQueryLimitError()
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		return nil, apperror.NewInvalidQueryPageError()
	}

	offset := (pageInt - 1) * limitInt
	return &dto.Query{
		Limit:  int64(limitInt),
		Page:   int64(pageInt),
		Offset: int64(offset),
		Sort:   sort,
		SortBy: sortBy,
		Search: query,
	}, nil
}
