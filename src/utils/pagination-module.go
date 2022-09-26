package utils

import (
	"database/sql"
	errors2 "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type PagingProcess interface {
	initialize(page, elemPerPage, paginateType int)
	DetachPagingParam(c *gin.Context)
	Processing(c *gin.Context, paginateType PaginateParam)
	GetTotalItemsCount(query *gorm.DB) (*PaginateDto, error)
}

type PaginateParam int64

const (
	LIMIT_OFFSET      PaginateParam = 0
	NESTED_OBJECT                   = 1
	DEFAULT_PAGE_SIZE               = 10
	DEFAULT_PAGE                    = 1
)

type PaginateDto struct {
	Page       int
	Size       int
	TotalItems int64
	TotalPages int64
}

type PaginateHelper struct {
	PagingProcess
	Page       int
	Limit      int
	Offset     int
	ElmPerPage int
	Start      int
	End        int
	PagingType PaginateParam
}

func (p *PaginateHelper) initialize(
	paginateType PaginateParam,
	DetachPagingParam func(c *gin.Context), c *gin.Context) {
	DetachPagingParam(c)
	fmt.Printf("%10s=%5d \n%10s=%5d\n", "Page", p.Page, "ElmPerPage", p.ElmPerPage)
	p.PagingType = paginateType
	if paginateType == LIMIT_OFFSET {
		p.Limit = p.ElmPerPage
		p.Offset = p.Limit + p.Page - 1
	}
	if paginateType == NESTED_OBJECT {
		p.Start = p.Page - 1
		p.End = p.ElmPerPage - 1
	}
}

func (p *PaginateHelper) GetTotalPages(TotalElements int) int {
	var totalPages int
	checkDivision := TotalElements % p.ElmPerPage
	if checkDivision == 0 {
		totalPages = TotalElements / p.ElmPerPage
	} else {
		totalPages = (TotalElements / p.ElmPerPage) + 1
	}
	return totalPages
}

func (p *PaginateHelper) GetTotalItemsCount(query *gorm.DB) (*PaginateDto, error) {
	var dbResult []map[string]interface{}
	RawSQLQueryScanRowsToMapHandler(query, &dbResult)
	PagingDto := &PaginateDto{}
	if query.RowsAffected > 0 {
		for _, data := range dbResult {
			if total, found := data["total"]; found && total != nil {
				PagingDto.TotalItems = total.(int64)
				PagingDto.TotalPages = int64(p.GetTotalPages(int(PagingDto.TotalItems)))
				PagingDto.Page = p.Page
				PagingDto.Size = p.ElmPerPage
				return PagingDto, nil
			}
		}
	}
	return nil, errors2.New("error while counting total")
}

func RawSQLQueryScanRowsToMapHandler(query *gorm.DB, source interface{}) {
	rows, err := query.Rows()

	if err != nil {
		fmt.Println(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	for rows.Next() {
		err := query.ScanRows(rows, source)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *PaginateHelper) DetachPagingParam() func(c *gin.Context) {
	return func(c *gin.Context) {
		p.ElmPerPage = convertQueryParam(c, "size", DEFAULT_PAGE_SIZE)
		p.Page = convertQueryParam(c, "page", DEFAULT_PAGE)
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
	p.initialize(paginateType, p.DetachPagingParam(), c)
	fmt.Printf("%10s=%5d \n%10s=%5d\n", "Page", p.Page, "ElmPerPage", p.ElmPerPage)

}

func (p *PaginateHelper) SetPagingParam(source, target interface{}) {

}
