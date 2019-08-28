package main

import (
	"strconv"
	"strings"

	"github.com/vjeantet/grok"
)

var grokPattern = "]%{GREEDYDATA:operation}: %{GREEDYDATA:filepath}:%{INT:lineNumber} %{GREEDYDATA:functionName}"

// Pattern matching using GROK
func matchPattern(input string) Result {
	g, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	values, err := g.Parse(grokPattern, input)
	if err != nil {
		return Result{}
	}
	var res Result
	for k, v := range values {
		// trimming the unnecassary stuff
		v = strings.Replace(v, " ", "", -1)
		v = strings.Replace(v, "\r", "", -1)

		if k == "operation" {
			res.Operation = v
		}
		if k == "filepath" {
			res.Filename = v
		}
		if k == "lineNumber" {
			num, _ := strconv.Atoi(v)
			res.LineNumber = num
		}
		if k == "functionName" {
			if v == "0" {
				res.Name = "anonymous"
				continue
			}
			res.Name = v
		}
	}
	return res
}

// Analysis : struct for complete analyzed response
type Analysis struct {
	Results []Result `json:"result"`
}

// Result : struct for one analyzed entity
type Result struct {
	Operation  string `json:"operation"`
	Filename   string `json:"filename"`
	LineNumber int    `json:"line_number"`
	Name       string `json:"name"`
}
