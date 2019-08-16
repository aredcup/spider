package models

import (
	"github.com/golang/glog"
	"redcup/spider/package/storage/realization/mysql/db"
	"redcup/spider/package/types"
	"strings"
)

type InsertResult interface {
	Insert(response types.Response)
}

type UserInfo struct {
	ID       uint64 `xorm:"id int(11) pk not null autoincr" description:"标识id"`
	NickName string `xorm:"nick_name varchar(45)  not null default ''" description:"昵称"`
	UserInfo string `xorm:"user_info varchar(255)  not null default ''" description:"信息"`
}

type Work struct {
	ID         uint32
	JobName    string
	Company    string
	Price      string
	Addr       string
	Experience string
	Education  string
	FullTime   string
	Welfare    string
}

func (m *Work) Insert(response types.Response) {
	for _, matchResult := range response.MatchRules {
		switch matchResult.Keys {
		case "job_name":
			m.Price = matchResult.Values
		case "price,addr,experience,education,full_time":
			values := strings.Split(matchResult.Values, ",")
			m.Price = values[0]
			m.Addr = values[1]
			m.Experience = values[2]
			m.Education = values[3]
			m.FullTime = values[4]
		case "company":
			m.Company = matchResult.Values
		case "welfare":
			m.Welfare = matchResult.Values
		}
	}

	_, err := db.GetWriteDB().Insert(m)
	if err != nil {
		glog.Error(err)
	}
}

func NewInsertResult() InsertResult {
	return &Work{}
}
