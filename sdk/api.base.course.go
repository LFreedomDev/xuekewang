package sdk

type Course struct {
	// 学科ID
	SubjectId int32 `json:"subject_id"`
	// 更新时间
	UpdateTime DateTime `json:"update_time"`
	// 创建时间
	CreateTime DateTime `json:"create_time"`
	// 学段ID
	StageId int32 `json:"stage_id"`
	// 课程名称
	Name string `json:"name"`
	// 课程ID
	Id int32 `json:"id"`
	// 排序值
	Ordinal int32 `json:"ordinal"`
}

// 获取课程列表
func (cli *SdkClient) GetCoursesAll() (res struct {
	ApiBaseResult
	Data []Course `json:"data"`
}, err error) {
	if err = cli.requestJSON("GET", "/xopqbm/courses/all", nil, nil, &res); err == nil {
		err = res.Error()
	}
	return
}

type GetCourseParams struct {
	// 课程ID
	// 	是否必须: true
	Id int32 `json:"id"`
}

// 获取指定ID的课程
func (cli *SdkClient) GetCourses(opts GetCourseParams) (res struct {
	ApiBaseResult
	Data Course `json:"data"`
}, err error) {
	if err = cli.requestJSON("GET", "/xopqbm/courses", ApiParams{"id": opts.Id}, nil, &res); err == nil {
		err = res.Error()
	}
	return
}
