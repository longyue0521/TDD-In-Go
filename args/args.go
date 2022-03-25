package args

import (
	"reflect"
	"strconv"
)

func Parse(v interface{}, flags ...string) {
	obj := reflect.ValueOf(v).Elem()
	for i, flag := range flags {
		if flag == "-l" {
			if obj.CanSet() && obj.Type().Field(0).IsExported() {
				obj.Field(0).SetBool(true)
				continue
			}
		} else if flag == "-p" {
			if obj.CanSet() && obj.Type().Field(1).IsExported() {
				p, _ := strconv.Atoi(flags[i+1])
				obj.Field(1).SetInt(int64(p))
				continue
			}
		} else if flag == "-d" {
			if obj.CanSet() && obj.Type().Field(2).IsExported() {
				obj.Field(2).SetString(flags[i+1])
				continue
			}
		}
	}
}


