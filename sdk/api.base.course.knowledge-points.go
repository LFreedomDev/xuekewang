package sdk

import "time"

// 获取指定课程的整棵知识树，适用于题库类应用
//
//	每个课程对应一棵知识树。
type GetCourseKnowledgePointsParams struct {
	// 课程ID
	// 	是否必须: true
	CourseId int32 `json:"course_id"`
}

type CourseKnowledgePoint struct {
	CourseId   int32     `json:"course_id"`   //	课程ID	integer(int32)
	UpdateTime time.Time `json:"update_time"` //	修改时间	string(date-time)
	Depth      int32     `json:"depth"`       //	节点深度，一级节点的深度为1，二级节点的深度为2，以此类推。	integer(int32)
	CreateTime time.Time `json:"create_time"` //	创建时间	string(date-time)
	ForLite    bool      `json:"for_lite"`    //	适用于精简版	boolean
	ParentId   int32     `json:"parent_id"`   //	父节点ID	integer(int32)
	Name       string    `json:"name"`        //	知识点名称	string
	RootId     int32     `json:"root_id"`     //	root节点的ID	integer(int32)
	Id         int32     `json:"id"`          //	知识点ID	integer(int32)
	Type       string    `json:"type"`        //	节点类型，可用值：NODE、KNOWLEDGE_POINT、TESTING_POINT，分别代表普通节点、知识点、考点。	string
	Ordinal    int32     `json:"ordinal"`     //	排序值	integer(int32)
}

// 获取指定课程的知识树
func (cli *SdkClient) GetCourseKnowledgePoints(opts GetCourseKnowledgePointsParams) (res struct {
	ApiBaseResult
	Data []CourseKnowledgePoint
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/courses/knowledge-points", NewApiParamsFromObject(opts), nil, &res)
	return
}
