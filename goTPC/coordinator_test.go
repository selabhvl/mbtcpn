package goTPC

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

var tpcTests XMLTest

func TestMain(m *testing.M) {
	// Flag definitions.
	var dir = flag.String(
		"dir",
		"xml/ss-tpctests-3.xml",
		"path to system test file",
	)

	// Parse and validate flags.
	flag.Parse()

	// Load the system test cases from XML file
	ParseXMLTestCase(*dir, &tpcTests)

	// Run tests/benchmarks.
	res := m.Run()
	os.Exit(res)
}

func TestNewCoordinator(t *testing.T) {

	t.Logf("The test cases from xml: %s", tpcTests.TestName)

	canCommitChannel := make(chan CanCommit, 16)
	decisionChannel := make(chan Decision, 16)
	finalDecisionChannel := make(chan DecisionEnum, 1)

	for _, cs := range tpcTests.TestCases {

		t.Logf("The test case ID: %v", cs.CaseID)
		t.Log("current number or workers:", cs.NumOfWorker)

		ch := NewCoordinator(canCommitChannel, decisionChannel)

		ch.Start(cs.NumOfWorker, finalDecisionChannel)

		for i := 0; i < cs.NumOfWorker; i++ {
			go func() {
				select {
				case c, ok := <-canCommitChannel:
					if !ok {
						t.Log("Haven't received cancommit message")
					} else {
						t.Log("Received cancommit message is:", c.String())
						idx := Search(cs.Votes, c, 0, len(cs.Votes))
						fmt.Println("got idx:", idx, "for worker:", c.GetWorkID())
						ch.DeliverVote(cs.Votes[idx])
					}
				}
			}()
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(20 * time.Millisecond)

		for j := 0; j < cs.NumOfWorker; j++ {
			go func() {
				select {
				case d, ok := <-decisionChannel:
					if !ok {
						t.Log("Haven't received decision message")
					} else {
						t.Log("Received decision message", d.String())
						idx := Search(cs.Decisions, d, 0, len(cs.Decisions))

						if d.GetDecision() != cs.Decisions[idx].GetDecision() {
							t.Errorf("got worker id=%d, decision=%s, expect worker id=%d, decision=%s", d.GetWorkID(), d.GetDecision(), cs.Decisions[idx].GetWorkID(), cs.Decisions[idx].GetDecision())
						}
						ch.DeliverACK(Ack{WorkerID: d.GetWorkID()})
					}
				}
			}()
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(40 * time.Millisecond)

		go func() {
			select {
			case fd, ok := <-finalDecisionChannel:
				if !ok {
					t.Log("Haven't received final decision")
				} else {
					t.Log("Received final decision from coordination:", fd.String())
					if fd != cs.FinalDecision {
						t.Errorf("got the final decision=%s from coordinator, expect final decision=%s", fd.String(), cs.FinalDecision.String())
					}
				}
			}
		}()
		time.Sleep(60 * time.Millisecond)

	}
	time.Sleep(80 * time.Millisecond)
}
