package utils

import (
	"delivery/constants"
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"time"
)

func PrintLog(err interface{}) {
	if err != nil {
		log.Print(err)
	}
}

func FormatStringTo24hourTime(datetime string) string {
	// parse to time
	t, err := time.Parse(constants.Kitchen, datetime) // parse with minute given
	if err != nil {
		t, err = time.Parse(constants.Kitchen2, datetime) // parse without minute given
		if err != nil {
			log.Print(err)
		}
	}
	return t.Format(constants.MilitaryTime)
}

func ReplaceStringRegex(regex string, text string, replaceText string) string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, replaceText)
}

func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		//use of IsNil method
		return reflect.ValueOf(val).IsNil()
	case reflect.Int, reflect.Float64, reflect.Uint, reflect.String:
		return false
	}
	if reflect.ValueOf(val).Kind() != reflect.Ptr && reflect.ValueOf(val).Len() == 0 {
		return true
	}
	return false
}

func AutoMap(from interface{}, to interface{}) error {
	jsonFrom, _ := json.Marshal(from)
	err := json.Unmarshal([]byte(string(jsonFrom)), to)
	return err
}

func ConvertEpochToTime(epochTime int) *time.Time {
	result := time.Unix(int64(epochTime), 0)
	return &result
}
