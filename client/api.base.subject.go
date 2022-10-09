package client

type Subject struct {
	// 学科名称
	Name string `json:"name"`
	// 学科ID
	Id int32 `json:"id"`
	// 排序值
	Ordinal int32 `json:"ordinal"`
}

// 获取学科列表
func (cli *SdkClient) GetSubjects() (res struct {
	ApiBaseResult
	Data []Subject
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/subjects", nil, nil, &res)
	return
}
