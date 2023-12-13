package db

import (
	"xorm.io/xorm"
)

type QueryFunc func() *xorm.Session

// Pagination 分页结构体
type Pagination struct {
	PageNo     int         // 当前页码
	PageSize   int         // 每页显示的数量
	TotalCount int64       // 总条数
	TotalPage  int         // 总页数
	PrevPage   int         // 上一页
	NextPage   int         // 下一页
	PageList   []int       // 页码列表
	Data       interface{} // 数据
}

// Paginate 分页函数，通过传入引擎和数据bean，实现对数据库表的分页查询
func (p *Pagination) Paginate(query QueryFunc, bean, data interface{}) error {

	// 使用xorm的Limit方法，根据PageNo和PageSize计算出Offset和Limit，实现对数据的切片查询
	err := query().Limit(p.PageSize, (p.PageNo-1)*p.PageSize).Find(data, bean)
	if err != nil {
		return err
	}

	p.Data = data
	// 获取记录总数
	totalCount, err := query().Count(bean)
	if err != nil {
		return err
	}

	if totalCount == 0 {
		return nil
	}

	// 计算所有相关的页码信息
	p.TotalCount = totalCount
	p.TotalPage = int(totalCount)/p.PageSize + 1
	p.PrevPage = p.PageNo - 1
	if p.PrevPage < 1 {
		p.PrevPage = 1
	}

	p.NextPage = p.PageNo + 1
	if p.NextPage > p.TotalPage {
		p.NextPage = p.TotalPage
	}

	// 调用页码列表生成函数
	p.PageList = p.generatePageList(p.PageNo, p.TotalPage, 5)
	return nil
}

// 页码列表生成函数
func (p *Pagination) generatePageList(currentPage, totalPages, displayCount int) []int {
	pages := make([]int, 0)

	// 计算起始和结束页码
	startPage := currentPage - (displayCount / 2)
	if startPage < 1 {
		startPage = 1
	}
	endPage := startPage + displayCount - 1
	if endPage > totalPages {
		endPage = totalPages
		startPage = endPage - displayCount + 1
		if startPage < 1 {
			startPage = 1
		}
	}

	// 创建分页对象列表
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}

	return pages
}

func (p *Pagination) FirstPage() int {
	return 1
}

func (p *Pagination) LastPage() int {
	lastPage := p.TotalCount / int64(p.PageSize)
	if p.TotalCount%int64(p.PageSize) != 0 {
		lastPage++
	}
	return int(lastPage)
}

// NewPagination 新建分页对象的函数，初始化当前页码和每页显示数量
func NewPagination(pageNo, pageSize int) *Pagination {
	return &Pagination{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
}
