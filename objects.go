package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	StoreObjectInDb(dbHdl redis.Conn) error
	DeleteObjectFromDb(dbHdl redis.Conn) error
	GetKey() string
	GetObjectFromDbByKey(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	GetObjectFromDb(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl redis.Conn) error
	GetAllObjFromDb(dbHdl redis.Conn) ([]ConfigObj, error)
	GetBulkObjFromDb(startIndex int64, count int64, dbHdl redis.Conn) (error, int64, int64, bool, []ConfigObj)
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler//

type User struct {
	ConfigObj
	UserName    string `SNAPROUTE: "KEY"`
	Password    string
	Description string
	Privilege   string
}

func (obj User) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userObj User
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userObj); err != nil {
			fmt.Println("### Trouble in unmarshaling User from Json", body)
		}
	}

	return userObj, err
}

type UserState struct {
	ConfigObj
	UserName      string `SNAPROUTE: "KEY"`
	LastLoginTime time.Time
	LastLoginIp   string
	NumAPICalled  uint32
}

func (obj UserState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userStateObj UserState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserState from Json", body)
		}
	}

	return userStateObj, err
}

type Login struct {
	ConfigObj
	UserName string
	Password string
}

func (obj Login) UnmarshalObject(body []byte) (ConfigObj, error) {
	var loginObj Login
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &loginObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Login from Json", body)
		}
	}

	return loginObj, err
}

type Logout struct {
	ConfigObj
	UserName  string
	SessionId uint32
}

func (obj Logout) UnmarshalObject(body []byte) (ConfigObj, error) {
	var logoutObj Logout
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &logoutObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Logout from Json", body)
		}
	}

	return logoutObj, err
}

type SystemStatusState struct {
	ConfigObj
	Name           string        `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Name of the system"`
	Ready          bool          `DESCRIPTION: "System is ready to accept api calls"`
	Reason         string        `DESCRIPTION: "Reason if system not ready"`
	UpTime         string        `DESCRIPTION: "Uptime of this system"`
	NumCreateCalls string        `DESCRIPTION: "Number of create api calls made"`
	NumDeleteCalls string        `DESCRIPTION: "Number of delete api calls made"`
	NumUpdateCalls string        `DESCRIPTION: "Number of update api calls made"`
	NumGetCalls    string        `DESCRIPTION: "Number of get api calls made"`
	NumActionCalls string        `DESCRIPTION: "Number of action api calls made"`
	FlexDaemons    []DaemonState `DESCRIPTION: "Daemon states"`
}

func (obj SystemStatusState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var systemStatus SystemStatusState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &systemStatus); err != nil {
			fmt.Println("### Trouble in unmarshaling SystemStatus from Json", body)
		}
	}

	return systemStatus, err
}

func (obj SystemStatusState) GetKey() string {
	return ""
}

type RepoInfo struct {
	Name   string `DESCRIPTION: "Name of the git repo"`
	Sha1   string `DESCRIPTION: "Git commit Sha1"`
	Branch string `DESCRIPTION: "Branch name"`
	Time   string `DESCRIPTION: "Build time"`
}

type SystemSwVersionState struct {
	ConfigObj
	FlexswitchVersion string     `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Flexswitch version"`
	Repos             []RepoInfo `DESCRIPTION: "Git repo details"`
}

func (obj SystemSwVersionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var systemSwVersion SystemSwVersionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &systemSwVersion); err != nil {
			fmt.Println("### Trouble in unmarshaling SystemSwVersion from Json", body)
		}
	}

	return systemSwVersion, err
}

func (obj SystemSwVersionState) GetKey() string {
	return ""
}
