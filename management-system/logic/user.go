package logic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/qiniu/qmgo/field"
)

type User struct {
	field.DefaultField `bson:",inline"`
	UserName           string `bson:"user_name" json:"userName"`
	Email              string `bson:"email" json:"email"`
	Password           string `bson:"password" json:"password"`
}

func (u *User) String() string {
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
