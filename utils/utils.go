package utils

import (
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

func ToInt64(s interface{}) int64 {
	i, _ := strconv.Atoi(s.(string))
	return int64(i)
}

func ToUint(s interface{}) uint {
	i, _ := strconv.Atoi(s.(string))
	return uint(i)
}

func StoBool(s interface{}) bool {
	return s.(string) == "true"
}

func ToInt(s interface{}) (int, error) {
	i, err := strconv.Atoi(s.(string))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func GetStartAndEndOfMonth() (time.Time, time.Time) {
	return now.BeginningOfMonth(), now.EndOfMonth()
}
