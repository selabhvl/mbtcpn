package goTPC

import (
	"encoding/xml"
	"io/ioutil"
)

// XMLTest defines variables used to map to xml format
type XMLTest struct {
	XMLName   xml.Name       `xml:"Test"`
	TestName  string         `xml:"TestName,attr"`
	TestCases []XMLTestCases `xml:"TestCase"`
}

// XMLTestCases defines variables for test cases mapped to xml format
type XMLTestCases struct {
	XMLName       xml.Name      `xml:"TestCase"`
	CaseID        string        `xml:"CaseID,attr"`
	NumOfWorker   int           `xml:"NumOfWorker,attr"`
	Votes         VoteSlice     `xml:"TestValues>Vote"`
	Decisions     DecisionSlice `xml:"TestOracles>Decision"`
	FinalDecision DecisionEnum  `xml:"TestOracles>FinalDecision"`
}

// ParseXMLTestCase read values from xml
func ParseXMLTestCase(file string, xmlTestCaseType interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, &xmlTestCaseType)
}
