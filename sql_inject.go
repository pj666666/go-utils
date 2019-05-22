package utils

import (
	"regexp"
)

//防SQL注入检查
//过滤 ‘
//ORACLE 注解 --  /**/
//关键字过滤 update,delete
// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
func CheckSQLInject(s string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		return false
	}
	return re.MatchString(s)
}
