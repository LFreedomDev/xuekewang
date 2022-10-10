package sdk

type Subject struct {
	// 学科名称
	Name string `json:"name"`
	// 学科ID
	Id int32 `json:"id"`
	// 排序值
	Ordinal int32 `json:"ordinal"`
}

// 获取学科列表
func (cli *SdkClient) GetSubjects(dataTemplate ...interface{}) (res struct {
	ApiBaseResult
	Data interface{} `json:"data"`
}, err error) {
	if dataTemplate == nil || len(dataTemplate) <= 0 {
		dataTemplate = []interface{}{[]Subject{}}
	}
	res.Data = dataTemplate[0]
	if err = cli.requestJSON("GET", "/xopqbm/subjects", nil, nil, &res); err == nil {
		err = res.Error()
	}
	return
}
