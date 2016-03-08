package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj BGPPolicyStmtConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPPolicyStmtConfig " +
		"( " +
		"Name TEXT, " +
		"MatchConditions TEXT, " +
		"PRIMARY KEY(Name) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)

	dbCmd = "CREATE TABLE IF NOT EXISTS BGPPolicyStmtCondition " +
		"( " +
		"Statement TEXT, " +
		"Condition TEXT, " +
		"FOREIGN KEY(Statement) REFERENCES BGPPolicyStmtConfig(Name) ON DELETE CASCADE, " +
		"FOREIGN KEY(Condition) REFERENCES BGPPolicyConditionConfig(Name) ON DELETE CASCADE, " +
		"PRIMARY KEY(Statement, Condition) " +
		")"

	_, err1 := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err == nil {
		err = err1
	}

	dbCmd = "CREATE TABLE IF NOT EXISTS BGPPolicyStmtAction " +
		"( " +
		"Statement TEXT, " +
		"Action TEXT, " +
		"FOREIGN KEY(Statement) REFERENCES BGPPolicyStmtConfig(Name) ON DELETE CASCADE, " +
		"FOREIGN KEY(Action) REFERENCES BGPPolicyActionConfig(Name) ON DELETE CASCADE, " +
		"PRIMARY KEY(Statement, Action) " +
		")"

	_, err2 := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err == nil {
		err = err2
	}
	return err
}

func (obj BGPPolicyStmtConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPPolicyStmtConfig (Name, MatchConditions) VALUES ('%v', '%v') ;",
		obj.Name, obj.MatchConditions)
	fmt.Println("**** Insert BGPPolicyStmtConfig called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to execute statement", dbCmd, "on BGPPolicyStmtConfig", err)
	} else {
		objectId, err = result.LastInsertId()
		if err != nil {
			fmt.Println("### Failed to return last object id", err)
		}
	}

	for conditionIdx := 0; conditionIdx < len(obj.Conditions); conditionIdx++ {
		dbCmd = fmt.Sprintf("INSERT INTO BGPPolicyStmtCondition (Statement, Condition) VALUES ('%v', '%v') ;",
			obj.Name, obj.Conditions[conditionIdx])
		fmt.Println("**** Insert BGPPolicyStmtCondition called with ", obj)

		result, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
		if err != nil {
			fmt.Println("**** Failed to execute statement", dbCmd, "on BGPPolicyStmtCondition", err)
		}
	}

	for actionIdx := 0; actionIdx < len(obj.Actions); actionIdx++ {
		dbCmd = fmt.Sprintf("INSERT INTO BGPPolicyStmtAction (Statement, Action) VALUES ('%v', '%v') ;",
			obj.Name, obj.Actions[actionIdx])
		fmt.Println("**** Insert BGPPolicyStmtAction called with ", obj)

		result, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
		if err != nil {
			fmt.Println("**** Failed to execute statement", dbCmd, "on BGPPolicyStmtAction", err)
		}
	}

	return objectId, err
}

func (obj BGPPolicyStmtConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPPolicyStmtConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from BGPPolicyStmtConfig where " + sqlKey
	fmt.Println("### DB Deleting BGPPolicyStmtConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPPolicyStmtConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object BGPPolicyStmtConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from BGPPolicyStmtConfig where " + sqlKey
	var tmp2 string
	var tmp3 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Name, &object.MatchConditions)
	fmt.Println("### DB Get BGPPolicyStmtConfig\n", err)

	if err == nil {
		dbCmd = "select * from BGPPolicyStmtCondition where STATEMENT=\"" + object.Name + "\""
		conditionRows, err1 := dbHdl.Query(dbCmd)
		if err1 == nil {
			defer conditionRows.Close()
			for conditionRows.Next() {
				if err1 = conditionRows.Scan(&tmp2, &tmp3); err1 == nil {
					object.Conditions = append(object.Conditions, tmp3)
				}
			}
		} else if err == nil {
			err = err1
		}

		dbCmd = "select * from BGPPolicyStmtAction where STATEMENT =\"" + object.Name + "\""
		actionRows, err2 := dbHdl.Query(dbCmd)
		if err2 == nil {
			defer actionRows.Close()
			for actionRows.Next() {
				if err2 = actionRows.Scan(&tmp2, &tmp3); err2 == nil {
					object.Actions = append(object.Actions, tmp3)
				}
			}
		} else if err == nil {
			err = err2
		}
	}

	return object, err
}

func (obj BGPPolicyStmtConfig) GetKey() (string, error) {
	key := string(obj.Name)
	return key, nil
}

func (obj BGPPolicyStmtConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Name = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *BGPPolicyStmtConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*BGPPolicyStmtConfig, e error) {
	dbCmd := "select * from BGPPolicyStmtConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'BGPPolicyStmtConfig' with error BGPPolicyStmtConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	var tmp2 string
	var tmp3 string
	for rows.Next() {

		object := new(BGPPolicyStmtConfig)
		if err = rows.Scan(&object.Name, &object.MatchConditions); err != nil {

			fmt.Println("Db method Scan failed when interating over BGPPolicyStmtConfig")
		}

		conditionCmd := "select * from BGPPolicyStmtCondition where STATEMENT =\"" + object.Name + "\""
		conditionRows, err := dbHdl.Query(conditionCmd)
		if err == nil {
			for conditionRows.Next() {
				if err = conditionRows.Scan(&tmp2, &tmp3); err == nil {
					object.Conditions = append(object.Conditions, tmp3)
				}
			}
			conditionRows.Close()
		}

		actionCmd := "select * from BGPPolicyStmtAction where STATEMENT =\"" + object.Name + "\""
		actionRows, err := dbHdl.Query(actionCmd)
		if err == nil {
			for actionRows.Next() {
				if err = actionRows.Scan(&tmp2, &tmp3); err == nil {
					object.Actions = append(object.Actions, tmp3)
				}
			}
			actionRows.Close()
		}
	}
	return objList, nil
}
func (obj BGPPolicyStmtConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(BGPPolicyStmtConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		fieldTyp := objTyp.Field(i)
		if fieldTyp.Anonymous {
			continue
		}

		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if _, ok := updateKeys[fieldTyp.Name]; ok {
			if objVal.Kind() == reflect.Int {
				if int(objVal.Int()) != int(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int8 {
				if int8(objVal.Int()) != int8(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int16 {
				if int16(objVal.Int()) != int16(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int32 {
				if int32(objVal.Int()) != int32(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int64 {
				if int64(objVal.Int()) != int64(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint {
				if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint8 {
				if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint16 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint32 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint64 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Bool {
				if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
					attrIds[idx] = true
				}
			} else {
				if objVal.String() != dbObjVal.String() {
					attrIds[idx] = true
				}
			}
			if attrIds[idx] {
				fmt.Println("attribute changed ", fieldTyp.Name)
			}
		}
		idx++

	}
	return attrIds[:idx], nil
}

func (obj BGPPolicyStmtConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedBGPPolicyStmtConfig BGPPolicyStmtConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedBGPPolicyStmtConfig)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[idx] {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 ||
				dbObjField.Kind() == reflect.Uint64 {
				mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 {
				mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
		idx++

	}
	return mergedBGPPolicyStmtConfig, nil
}

func (obj BGPPolicyStmtConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbBGPPolicyStmtConfig := dbObj.(BGPPolicyStmtConfig)
	objKey, err := dbBGPPolicyStmtConfig.GetKey()
	objSqlKey, err := dbBGPPolicyStmtConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "BGPPolicyStmtConfig" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		if attrSet[idx] {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int ||
				fieldVal.Kind() == reflect.Int8 ||
				fieldVal.Kind() == reflect.Int16 ||
				fieldVal.Kind() == reflect.Int32 ||
				fieldVal.Kind() == reflect.Int64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else if fieldVal.Kind() == reflect.Uint ||
				fieldVal.Kind() == reflect.Uint8 ||
				fieldVal.Kind() == reflect.Uint16 ||
				fieldVal.Kind() == reflect.Uint32 ||
				fieldVal.Kind() == reflect.Uint64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
			} else if fieldVal.Kind() == reflect.Bool {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
		idx++
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
