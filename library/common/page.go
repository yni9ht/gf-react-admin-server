package common

// 分页统一返回体
type PageResult struct {
	PageNum  int         `json:"pageNum"`  // 分页页数
	PageSize int         `json:"pageSize"` // 分页大小
	Total    int         `json:"total"`    // 总记录数
	Records  interface{} `json:"records"`  // 对应数据
}

// NewResult init page result. 创建一个分页返回体
func NewResult(pageNum, pageSize, total int, records interface{}) PageResult {
	return PageResult{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		Records:  records,
	}
}
