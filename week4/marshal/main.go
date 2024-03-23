package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type stu struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func Marshal(a interface{}) ([]byte, error) {
	val := reflect.ValueOf(a)
	typ := val.Type()

	//检查一下是否是指针
	if typ.Kind() == reflect.Ptr {
		if val.IsNil() {
			return []byte("null"), nil
		} else {
			typ = typ.Elem()
			val = val.Elem()
		}
	}

	res := bytes.Buffer{}
	switch typ.Kind() {
	case reflect.String: //字符串
		return []byte(fmt.Sprintf("\"%s\"", val.String())), nil
	case reflect.Bool: //布尔
		return []byte(fmt.Sprintf("%t", val.Bool())), nil
	case reflect.Float32, reflect.Float64: //浮点数
		return []byte(fmt.Sprintf("%f", val.Float())), nil
	case reflect.Int: //整型
		return []byte(fmt.Sprintf("%d", val.Int())), nil
	case reflect.Slice: //切片
		if val.IsNil() {
			return []byte("null"), nil
		}
		res.WriteByte('[')
		if val.Len() > 0 {
			for i := 0; i < val.Len(); i++ {
				if bs, err := Marshal(val.Index(i).Interface()); err != nil {
					return nil, err
				} else {
					res.Write(bs)
					if i != val.Len()-1 {
						res.WriteByte(',')
					}

				}
			}
		}
		res.WriteByte(']')
		return res.Bytes(), nil
	case reflect.Map: //map
		res.WriteByte('{')
		if val.Len() > 0 {
			for _, key := range val.MapKeys() {
				if keybs, err := Marshal(key.Interface()); err != nil {
					return nil, err
				} else {
					res.Write(keybs)
					res.WriteByte(':')
					if vbs, err := Marshal(val.MapIndex(key).Interface()); err != nil {
						return nil, err
					} else {
						res.Write(vbs)
						res.WriteByte(',')
					}
				}
			}
			res.Truncate(len(res.Bytes()) - 1)
		}
		res.WriteByte('}')
		return res.Bytes(), nil
	case reflect.Struct: //结构体
		res.WriteByte('{')
		if val.NumField() > 0 {
			for i := 0; i < val.NumField(); i++ {
				fieldval := val.Field(i)
				fieldtyp := typ.Field(i)
				name := fieldtyp.Name
				if len(fieldtyp.Tag.Get("json")) > 0 {
					name = fieldtyp.Tag.Get("json")
				}
				res.WriteString(fmt.Sprintf("\"%s\":", name))
				if bs, err := Marshal(fieldval.Interface()); err != nil { //对val递归调用Marshal序列化
					return nil, err
				} else {
					res.Write(bs)
				}
				if i != val.NumField()-1 {
					res.WriteByte(',')
				}
			}
		}
		res.WriteByte('}')
		return res.Bytes(), nil
	default:
		return []byte(fmt.Sprintf("\"暂不支持该数据类型:%s\"", typ.Kind().String())), errors.New(fmt.Sprintf("类型未知"))
	}
}
func main() {
	sli := []stu{{ID: 1, Username: "陈诚"}, {ID: 2, Username: "小明"}}
	res, err := Marshal(sli)
	if err != nil {
		println(err)
	}
	fmt.Println(string(res))
}
