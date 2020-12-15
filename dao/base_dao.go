/**
 * @Title  数据库操作层
 * @Description 数据库操作封装的方法
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:08
 */
package dao

import (
	strings2 "github.com/ganeryao/linking-go-agile/strs"
	"reflect"
	"strings"
)

type BaseDao struct {
}

func marshalField(result *strings2.StringBuilder, fList *strings2.StringBuilder,
	vList *strings2.StringBuilder, obj interface{}, args []interface{}) []interface{} {
	objType := reflect.TypeOf(obj)
	var types reflect.Type
	var values reflect.Value
	if objType.Kind() == reflect.Struct {
		types = reflect.TypeOf(obj)
		values = reflect.ValueOf(obj)
	} else {
		types = reflect.TypeOf(obj).Elem()
		values = reflect.ValueOf(obj).Elem()
	}
fieldFor:
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		value := values.Field(i)
		if field.Type.Kind() == reflect.Struct {
			args = marshalField(result, fList, vList, value.Interface(), args)
		} else {
			dbTag := field.Tag.Get("db")
			tagArr := strings.Split(dbTag, ",")
			name := field.Name
			isKey := false
			for _, v := range tagArr {
				switch v {
				case "":
					continue
				case "key":
					isKey = true
				case "select":
				case "-":
					continue fieldFor
				default:
					name = v
				}
			}
			args = append(args, value.Interface())
			if !fList.IsEmpty() {
				result.Append(",")
				fList.Append(",")
			}
			result.Append(name)
			fList.Append("?")
			if !isKey {
				if !vList.IsEmpty() {
					vList.Append(",")
				}
				vList.Append(name)
				vList.Append("=values(")
				vList.Append(name)
				vList.Append(")")
			}
		}
	}
	return args
}

func (dao *BaseDao) MarshalUpSql(v interface{}, table string) (string, []interface{}) {
	result := strings2.NewStringBuilder()
	result.Append("INSERT INTO ")
	result.Append(table)
	result.Append("(")
	fields := strings2.NewStringBuilder()
	values := strings2.NewStringBuilder()
	var args []interface{}
	args = marshalField(result, fields, values, v, args)
	result.Append(") VALUES(")
	result.Append(fields.ToString())
	result.Append(") ON DUPLICATE KEY UPDATE ")
	result.Append(values.ToString())
	result.Append(";")
	return result.ToString(), args
}

func (dao *BaseDao) MarshalFieldSql(v interface{}, table string) (string, []interface{}) {
	result := strings2.NewStringBuilder()
	result.Append("select ")
	fields := strings2.NewStringBuilder()
	values := strings2.NewStringBuilder()
	var args []interface{}
	args = marshalField(result, fields, values, v, args)
	result.Append(" from ")
	result.Append(table)
	result.Append(" ")
	return result.ToString(), args
}

func (dao *BaseDao) MarshalSelectByPrimarySql(v interface{}, table string) (string, []interface{}) {
	result := strings2.NewStringBuilder()
	result.Append("select ")
	fields := strings2.NewStringBuilder()
	values := strings2.NewStringBuilder()
	var args []interface{}
	args = marshalField(result, fields, values, v, args)
	result.Append(" from ")
	result.Append(table)
	result.Append("  where id = ? limit 1")
	return result.ToString(), args
}
