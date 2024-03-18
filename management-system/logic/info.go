package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/qiniu/qmgo/field"
)

// var MongodbClient *qmgo.Client

// func InitMongodb(addr string) {

// 	ctx := context.Background()
// 	c, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: addr})
// 	if err != nil {
// 		fmt.Printf("err.Error(): %v\n", err.Error())
// 		return
// 	}
// 	MongodbClient = c
// }

type BaseInfo struct {
	field.DefaultField        `bson:",inline"`
	Manager                   string `bson:"manager" json:"manager"`
	LabsLeads                 string `bson:"labsLeads" json:"labsLeads"`
	PmCioName                 string `bson:"pmCioName" json:"pmCioName"`
	Type                      string `bson:"type" josn:"type"`
	FormerEmployer            string `bson:"formerEmployer" json:"formerEmployer"`
	FirmAUM                   string `bson:"firmAUM" json:"firmAUM"`
	Source                    string `bson:"source" json:"source"`
	AvailableStrategyCapacity string `bson:"asCapacity" json:"asCapacity"`
	LinkedinPage              string `bson:"linkedinPage" json:"linkedinPage"`
	EmailContact              string `bson:"emailContact" json:"emailContact"`
	FilePath                  string `bson:"filePath" json:"filePath"`
	CurrentStage              string `json:"currentStage" bson:"currentStage"`
	NeedRevisit               string `json:"needRevisit" bson:"needRevisit"`
	HowRevisit                string `json:"howRevisit" bson:"howRevisit"`
	CreateBy                  string `json:"createBy" bson:"createBy"`
}

type CommHistory struct {
	ReviewDate         time.Time `json:"reviewDate" bson:"reviewDate"`
	field.DefaultField `bson:",inline"`
	BaseInfoId         string `json:"baseInfoId" bson:"baseInfoId"`
	CurrentStage       string `json:"currentStage" bson:"currentStage"`
	Comment            string `json:"comment" bson:"comment"`
	InvestConviction   string `json:"investConviction" bson:"investConviction"`
	IsRevisit          string `json:"isRevisit" bson:"isRevisit"`
	HowRevisit         string `json:"howRevisit" bson:"howRevisit"`
}

func (u *BaseInfo) String() string {
	b, err := json.Marshal(*u)
	if err != nil {
		return fmt.Sprintf("%+v", *u)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *u)
	}
	return out.String()
}
