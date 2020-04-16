package goamazon

import (
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/expert"
	"strings"
)

func Is404(content []byte) bool {
	doc, _ := expert.QueryBytes(content)
	text := doc.Find("title").Text()
	if strings.Contains(text, "Page Not Found") {
		return true
	}
	if strings.Contains(text, "404") {
		return true
	}
	//uk
	if strings.Contains(string(content), "The Web address you entered is not a functioning page on our site") {
		return true
	}
	//de
	if strings.Contains(string(content), "Suchen Sie bestimmte Informationen") {
		return true
	}
	if strings.Contains(string(content), "Suchen Sie etwas bestimmtes") {
		return true
	}

	if text == "" {
		return true
	}
	return false
}

func IsRobot(content []byte) (s string) {
	doc, _ := expert.QueryBytes(content)
	text := doc.Find("title").Text()

	if strings.Contains(text, "Sorry! Something went wrong!") {
		return "sorry"
	}

	// uk usa
	if strings.Contains(text, "Robot Check") {
		return "robot"
	}
	//jp
	if strings.Contains(text, "CAPTCHA") {
		return "robot"
	}
	//de
	if strings.Contains(text, "Bot Check") {
		return "robot"
	}
	return
}

func TooSortSizes(data []byte, sizes float64) error {
	if float64(len(data))/1000 < sizes {
		return errors.New(fmt.Sprintf("FileSize:%d bytes,%d kb < %f kb dead too sort", len(data), len(data)/1000, sizes))
	}
	return nil
}
