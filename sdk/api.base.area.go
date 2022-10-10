package sdk

import (
	"time"
)

type Area struct {
	Code       string    `json:"code"`        // 行政区标准代码(会随着国家行政区划调整而变化，不建议作为主键使用)
	CreateTime time.Time `json:"create_time"` // 创建时间
	Level      string    `json:"level"`       // 行政区级别 可用值: PROVINCE、CITY、COUNTY 分别代表省级、地市级、区县级
	ParentId   string    `json:"parent_id"`   // 所属行政区ID
	Name       string    `json:"name"`        // 行政区名称
	ShortName  string    `json:"short_name"`  // 行政区简称
	Id         string    `json:"id"`          // 行政区ID(含数字和字母。一旦创建是不会变化的，作为主键使用)
	Ordinal    int32     `json:"ordinal"`     // 排序值
}

// 获取行政区列表
//
//	提供行政区、行政区级别和学科网定制行政区（比如全国一、全国二等）的数据查询接口。
//	行政区数据以树形结构存储，分为3个级别，省(直辖市)/地市/区县。其中省市县的编码为6位。比如北京市的编码为110000。
//	特别说明：直辖市是省级行政区，为了和省级行政区对齐，市级仍然存储直辖市名称。举例：北京（直辖市）->北京（地市）->东城（区县）。
func (cli *SdkClient) GetAreasAll() (res struct {
	ApiBaseResult
	Data []Area
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/areas/all", nil, nil, &res)
	return
}

type GetAreasParams struct {
	// 行政区ID
	// 	是否必须: true
	Id string `json:"id"`
}

// 获取指定ID的行政区
func (cli *SdkClient) GetAreas(opts GetAreasParams) (res struct {
	ApiBaseResult
	Data []Area
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/areas", ApiParams{"id": opts.Id}, nil, &res)
	return
}
