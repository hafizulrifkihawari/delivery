package utils

import (
	"delivery/constants"
	"log"
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
