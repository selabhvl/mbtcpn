package reader

import (
	"encoding/xml"
	tpc "github.com/selabhvl/mbtcpn/goTPC/twophasecommit"
	"io/ioutil"
)

type XMLTest struct {
	XMLName   xml.Name       `xml:"Test"`
	TestName  string         `xml:"TestName,attr"`
	TestCases []XMLTestCases `xml:"TestCase"`
}

type XMLTestCases struct {
	XMLName       xml.Name          `xml:"TestCase"`
	CaseID        string            `xml:"CaseID,attr"`
	NumOfWorker   int               `xml:"NumOfWorker,attr"`
	Votes         tpc.VoteSlice     `xml:"TestValues>Votes"`
	Decisions     tpc.DecisionSlice `xml:"TestOracles>Decisions"`
	FinalDecision tpc.DecisionEnum  `xml:"TestOracles>FinalDecision"`
}

func ParseXMLTestCase(file string, xmlTestCaseType interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, &xmlTestCaseType)
}
