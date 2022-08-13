package main

import (
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	data := clipboard.Read(clipboard.FmtText)

	result := strings.ReplaceAll(string(data), "ue", "ü")
	result = strings.ReplaceAll(result, "Ue", "Ü")
	result = strings.ReplaceAll(result, "ae", "ä")
	result = strings.ReplaceAll(result, "Ae", "Ä")
	result = strings.ReplaceAll(result, "oe", "ö")
	result = strings.ReplaceAll(result, "Oe", "Ö")
	result = strings.ReplaceAll(result, "sss", "ß")

	clipboard.Write(clipboard.FmtText, []byte(result))
}
