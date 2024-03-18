package vo

import (
	"time"

	"github.com/qiniu/qmgo/field"
)

type PageRequest[T any] struct {
	PageNum  int64 ` form:"pageNum" json:"pageNum,omitempty"`
	PageSize int64 `form:"pageSize" json:"pageSize,omitempty"`
	Filter   T     `pform:"filter" json:"filter,omitempty"`
}

type ReqId struct {
	Id string `form:"id" json:"id" uri:"id"`
}

type ReqIds struct {
	Ids []string `form:"ids" json:"ids"`
}

type InfoVo struct {
	field.DefaultField        `bson:",inline"`
	Manager                   string `json:"manager" bson:"manager"`
	LabsLeads                 string `json:"labsLeads" bson:"labsLeads"`
	PmCioName                 string `json:"pmCioName" bson:"pmCioName"`
	Type                      string `json:"type" bson:"type"`
	FormerEmployer            string `json:"formerEmployer" bson:"formerEmployer"`
	FirmAUM                   string `json:"firmAUM" bson:"firmAUM"`
	Source                    string `json:"source" bson:"source"`
	AvailableStrategyCapacity string `json:"asCapacity" bson:"asCapacity"`
	LinkedinPage              string `json:"linkedinPage" bson:"linkedinPage"`
	EmailContact              string `json:"emailContact" bson:"emailContact"`
	FilePath                  string `json:"filePath" bson:"filePath"`
	// 回访信息
	ReviewDate       time.Time `json:"reviewDate" bson:"reviewDate"`
	BaseInfoId       string    `json:"baseInfoId" bson:"baseInfoId"`
	CurrentStage     string    `json:"currentStage" bson:"currentStage"`
	Comment          string    `json:"comment" bson:"comment"`
	InvestConviction string    `json:"investConviction" bson:"investConviction"`
	IsRevisit        string    `json:"isRevisit" bson:"isRevisit"`
	HowRevisit       string    `json:"howRevisit" bson:"howRevisit"`
}
