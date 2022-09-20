package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PagingProcess interface {
	initialize(page, elemPerPage, paginateType int)
	DetachPagingParam(c *gin.Context)
	Processing(c *gin.Context, paginateType PaginateParam)
}

type PaginateParam int64

const (
	LIMIT_OFFSET      PaginateParam = 0
	NESTED_OBJECT                   = 1
	DEFAULT_PAGE_SIZE               = 10
	DEFAULT_PAGE                    = 1
)

type PaginateHelper struct {
	PagingProcess
	page       int
	limit      int
	offset     int
	elmPerPage int
	start      int
	end        int
	pagingType PaginateParam
}

func (p *PaginateHelper) initialize(page int, elemPerPage int,
	paginateType PaginateParam,
	DetachPagingParam func(c *gin.Context), c *gin.Context) {
	DetachPagingParam(c)
	p.pagingType = paginateType
	p.page = page
	p.elmPerPage = elemPerPage
	if paginateType == LIMIT_OFFSET {
		p.limit = elemPerPage
		p.offset = p.limit + p.page - 1
	}
	if paginateType == NESTED_OBJECT {
		p.start = p.page - 1
		p.end = p.elmPerPage - 1
	}
}

func (p *PaginateHelper) DetachPagingParam() func(c *gin.Context) {
	return func(c *gin.Context) {
		p.elmPerPage = convertQueryParam(c, "size", DEFAULT_PAGE_SIZE)
		p.page = convertQueryParam(c, "page", DEFAULT_PAGE)
	}

}

func convertQueryParam(c *gin.Context, name string, defaultValue int) int {
	Cvt := c.Query(name)
	cvt, err := strconv.Atoi(Cvt)
	if err != nil {
		return defaultValue
	}
	return cvt
}

func (p *PaginateHelper) Processing(c *gin.Context, paginateType PaginateParam) {
	p.initialize(p.page, p.elmPerPage, paginateType, p.DetachPagingParam(), c)
}
