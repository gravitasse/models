package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	CreateDBTable(dbHdl *sql.DB) error
	StoreObjectInDb(dbHdl *sql.DB) (int64, error)
	DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error
	GetKey() (string, error)
	GetSqlKeyStr(string) (string, error)
	GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl *sql.DB) error
	GetAllObjFromDb(dbHdl *sql.DB) ([]ConfigObj, error)
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler//

/* Start asicd owned objects */
/*
 * Vlan object and Route objects are exception cases, i.e
 * they are not seperated into config and state, sub-objects.
 * This approach is followed for objects that can be both
 * statically and dynamically created.
 */
type Vlan struct {
	BaseObj
	VlanId           int32 `SNAPROUTE: "KEY"`
	VlanName         string
	OperState        string
	IfIndex          int32
	IfIndexList      string
	UntagIfIndexList string
}	
type LogicalIntfConfig struct {
	BaseObj
	Name string `SNAPROUTE: "KEY"`
	Type string
}

func (obj LogicalIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf LogicalIntfConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### LogicalIntfConfig called, unmarshal failed", body)
		}
	}
	return gConf, err
}

type LogicalIntfState struct {
	BaseObj
	Name              string `SNAPROUTE: "KEY"`
	IfIndex           int32
	SrcMac            string
	OperState         string
	IfInOctets        int64
	IfInUcastPkts     int64
	IfInDiscards      int64
	IfInErrors        int64
	IfInUnknownProtos int64
	IfOutOctets       int64
	IfOutUcastPkts    int64
	IfOutDiscards     int64
	IfOutErrors       int64
}

func (obj LogicalIntfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf LogicalIntfState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling LogicalIntfState from Json", body)
		}
	}
	return gConf, err
}

type IPv4Intf struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY"`
	IfIndex int32
}

type PortConfig struct {
	BaseObj
	PortNum     int32 `SNAPROUTE: "KEY"`
	Description string
	PhyIntfType string
	AdminState  string
	MacAddr     string
	Speed       int32
	Duplex      string
	Autoneg     string
	MediaType   string
	Mtu         int32
}

type PortState struct {
	BaseObj
	PortNum           int32 `SNAPROUTE: "KEY"`
	IfIndex           int32
	Name              string
	OperState         string
	IfInOctets        int64
	IfInUcastPkts     int64
	IfInDiscards      int64
	IfInErrors        int64
	IfInUnknownProtos int64
	IfOutOctets       int64
	IfOutUcastPkts    int64
	IfOutDiscards     int64
	IfOutErrors       int64
	ErrDisableReason  string
}

func (obj PortState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf PortState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling PortState from Json", body)
		}
	}
	return gConf, err
}

/* End asicd owned objects */

type UserConfig struct {
	BaseObj
	UserName    string `SNAPROUTE: "KEY"`
	Password    string
	Description string
	Previledge  string
}

func (obj UserConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userConfigObj UserConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserConfig from Json", body)
		}
	}

	return userConfigObj, err
}

type UserState struct {
	BaseObj
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
	BaseObj
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
	BaseObj
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
