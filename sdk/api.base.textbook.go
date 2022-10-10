package sdk

type GetTextBooksParams struct {
	// 课程ID
	// 	是否必须: false
	CourseId string `json:"course_id"`
	// 年级ID
	// 	是否必须: false
	GradeId string `json:"grade_id"`
	// 当前页码（从1开始）,示例值(1)
	// 	是否必须: false
	PageIndex string `json:"page_index"`
	// 每页数据条数,示例值(50)
	// 	是否必须: false
	PageSize string `json:"page_size"`
	// 教材版本ID
	// 	是否必须: false
	VersionId string `json:"version_id"`
}

// 获取教材列表
func (cli *SdkClient) GetTextBooks(opts GetTextBooksParams) (res struct {
	ApiBaseResult
	Data interface{}
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/textbooks", ApiParams{
		"course_id":  opts.CourseId,
		"grade_id":   opts.GradeId,
		"page_index": opts.PageIndex,
		"page_size":  opts.PageSize,
		"version_id": opts.VersionId,
	}, nil, &res)
	return
}
