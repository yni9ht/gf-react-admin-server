package common

type PageReq struct {
	PageNo   int `json:"pageNo" d:"1" v:"required|min:1#分页页数不能为空|分页页数参数非法"`    // 分页页数
	PageSize int `json:"pageSize" d:"10" v:"required|min:1#分页大小不能为空|分页大小参数非法"` // 分页大小
}

// 分页统一返回体
type PageResult struct {
	PageNo   int         `json:"pageNo"`   // 分页页数
	PageSize int         `json:"pageSize"` // 分页大小
	Total    int         `json:"total"`    // 总记录数
	Records  interface{} `json:"records"`  // 对应数据
}

// NewResult init page result. 创建一个分页返回体
func NewResult(pageNo, pageSize, total int, records interface{}) PageResult {
	return PageResult{
		PageNo:   pageNo,
		PageSize: pageSize,
		Total:    total,
		Records:  records,
	}
}
