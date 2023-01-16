package service

import "{{.ProjectName}}/pkg/zaplogger"

type Users struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Hobby []string `json:"hobby"`
}

// GetUserInfoService 获取用户信息
func GetUserInfoService() (userList []Users, err error) {
	// 1.组装数据
	userList = []Users{
	    {
	        ID: 1,
	        Name: "张三",
	        Age: 19,
	        Hobby: []string{"看书", "听歌"},
	    },
	    {
	        ID: 2,
            Name: "李四",
            Age: 22,
            Hobby: []string{"跑步", "睡觉"},
        },
	}

    zaplogger.SugarLogger.Errorf("userList:%v\n", userList)
	// 2. 返回数据
	return userList, nil
}