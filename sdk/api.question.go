package sdk

import (
	"time"
)

// 章节知识点挑题接口请求参数（JSON格式），最大允许返回前100条数据
type QuestionPickParams struct {
	CourseId         int32   `json:"course_id"`          // 课程ID   是否必须:false 数据类型:integer(int32)
	AreaIds          []int32 `json:"area_ids"`           // 行政区ID列表（支持按省份、县市进行搜索），最多传10个  是否必须:false 数据类型:array
	FormulaPicFormat string  `json:"formula_pic_format"` // 公式图片格式，支持两种：png或svg，默认是svg  是否必须:false 数据类型:string
	Year             int32   `json:"year"`               // 年份（查询此年份及以后的试题）  是否必须:false integer(int32)
	KpointIds        []int32 `json:"kpoint_ids"`         // 试题知识点ID集合，最多传10个；如果传知识点父节点，也会搜索出其子节点中的试题  是否必须:false 数据类型:array
	TypeIds          []int32 `json:"type_ids"`           // 试题类型ID集合，最多传10个；如果传试题类型父节点，也会搜索出其子节点中的试题  是否必须:false 数据类型:array
	PageIndex        int32   `json:"page_index"`         // 当前页码（从1开始）  是否必须:false integer(int32)
	OrderBy          int32   `json:"order_by"`           // 排序(默认按照发布时间)：0=发布时间（倒序）,1=试题新鲜度（升序）  是否必须:false 数据类型:integer(int32)
	PaperTypeIds     []int32 `json:"paper_type_ids"`     // 试卷类型ID集合，最多传10个  是否必须:false 数据类型:array
	CatalogIds       []int32 `json:"catalog_ids"`        // 章节ID集合，最多传10个；如果传章节父节点，也会搜索出子节点中的试题  是否必须:false 数据类型:array
	DifficultyLevels []int32 `json:"difficulty_levels"`  // 试题难度等级ID集合（17 容易 18 较易 19 一般 20 较难 21 困难），最多传5个  是否必须:false 数据类型:array
	PageSize         int32   `json:"page_size"`          // 每页数据条数  是否必须:false 数据类型:integer(int32)
}

type QuestionPickData struct {
	PageIndex int32 `json:"page_index"` //	当前页码（从1开始）	integer(int32)
	PageSize  int32 `json:"page_size"`  //	每页记录数	integer(int32)
	TotalPage int32 `json:"total_page"` //	总页数	integer(int32)
	TotalSize int32 `json:"total_size"` //	总记录数	integer(int32)
	Items     []struct {
		CourseId        int32   `json:"course_id"`        //	课程ID	integer(int32)
		DifficultyLevel int32   `json:"difficulty_level"` //	试题难度等级（17 容易 18 较易 19 一般 20 较难 21 困难）	integer(int32)
		KpointIds       []int32 `json:"kpoint_ids"`       //	知识点ID列表	array
		TypeId          string  `json:"type_id"`          //	试题类型ID	string
		DynamicId       string  `json:"dynamic_id"`       //	动态试题ID	string
		Kpoints         []struct {
			Name string `json:"name"` //	名称	string
			Id   int    `json:"id"`   //	ID	integer
		} `json:"kpoints"` //	知识点列表	array
		PaperTypeIds []int32 `json:"paper_type_ids"` //	试卷类型ID列表	array
		Type         struct {
			Name string `json:"name"` //	名称	string
			Id   string `json:"id"`   //	ID	string
		} `json:"type"` //	试题类型	IdNamePair«string»

		CatalogIds   []int32 `json:"catalog_ids"` //	教材目录ID列表	array
		Years        []int32 `json:"years"`       //	试题出现在试卷中的年份，可能多个	array
		SourcePapers []struct {
			Area struct {
				Name string `json:"name"` //	名称	string
				Id   string `json:"id"`   //	ID	string
			} `json:"area"` //	地区	IdNamePair«string»
			Title string `json:"title"` //	试卷名称	string

		} `json:"source_papers"` //	来源试卷所在地区列表	array

		Difficulty int32 `json:"difficulty"` //	试题难度，0~1之间的数字，值越小难度越大（(0.9,1] 容易，(0.8,0.9] 较易，(0.5,0.8] 一般，(0.3,0.5] 较难，[0, 0.3] 困难）	number(double)
		Catalogs   []struct {
			Name string `json:"name"` //	名称	string
			Id   int    `json:"id"`   //	ID	integer
		} `json:"catalogs"` //	教材目录列表	array

		Course struct {
			Name string `json:"name"` //	名称	string
			Id   int    `json:"id"`   //	ID	integer
		} `json:"course"` //	课程	IdNamePair«int»

		AnswerScoreable int32     `json:"answer_scoreable"` //	在线作答，0=不支持，1=支持。选择题或者打标了机阅的试题	integer(int32)
		CreateDate      time.Time `json:"create_date"`      //	试题入库日期	string(date-time)
		Stem            string    `json:"stem"`             //	试题题干（HTML格式），请参考《试题结构和HTML渲染说明文档》	string
	} `json:"items"` //	本页数据	array
}

// 章节-知识点挑题
func (cli *SdkClient) QuestionPick(opts QuestionPickParams) (res struct {
	ApiBaseResult
	Data QuestionPickData
}, err error) {
	err = cli.requestJSON("POST", "/xopqbm/questions/pick", nil, NewApiParamsFromObject(opts), &res)
	return
}
