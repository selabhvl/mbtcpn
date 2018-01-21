package twophasecommit

import "fmt"

// The number of workers.
const numOfWorker = 3

// WorkerID represents a identifier for workers.
type WorkerID int

// VoteEnum represents the enumeration type for the value of the vote.
type VoteEnum int

const (
	Yes VoteEnum = iota
	No
)

var voteValues = []string{"Yes", "No"}

// DecisionEnum represents the enumeration type for the value of the decision.
type DecisionEnum int

const (
	Commit DecisionEnum = iota
	Abort
)

var decisionValues = []string{"Commit", "Abort"}

// ConCommit represents the message sent by the coordinator to the test adapter/workers.
type CanCommit struct {
	WorkerID WorkerID
}

// Vote represents the message sent by the test adapter/workers to the coordinator.
type Vote struct {
	WorkerID  WorkerID
	VoteValue VoteEnum
}

// Decision represents the message sent by the coordinator to the test adapter/workers..
type Decision struct {
	WorkerID      WorkerID
	DecisionValue DecisionEnum
}

// ACK represents the message sent by the test adapter/workers to the coordinator.
type Ack struct {
	WorkerID WorkerID
}

// String returns a string representation of the value of VoteEnum ve.
func (ve VoteEnum) String() string {
	return voteValues[ve]
}

// String returns a string representation of the value of DecisionEnum de.
func (de DecisionEnum) String() string {
	return decisionValues[de]
}

// String returns a string representation of cancommit c.
func (cc CanCommit) String() string {
	return fmt.Sprintf("CanCommit{worker: %d}", cc.WorkerID)
}

// String returns a string representation of vote v.
func (v Vote) String() string {
	return fmt.Sprintf("Vote{worker: %d, vote: %s}", v.WorkerID, v.VoteValue)
}

// String returns a string representation of decision d.
func (d Decision) String() string {
	return fmt.Sprintf("Decision{worker: %d, decision: %s}", d.WorkerID, d.DecisionValue)
}

// String returns a string representation of ack a.
func (a Ack) String() string {
	return fmt.Sprintf("ACK{worker: %d}", a.WorkerID)
}

// Coordinator represent a coordinator for the protocol.
type Coordinator interface {
	Start(numOfWorker int, finalDecisionChannel chan DecisionEnum)
	SendCanCommit(cc CanCommit)
	DeliverVote(v Vote)
	CollectVotes(v Vote, votes []Vote) []Vote
	SendDecision(d Decision)
	DeliverACK(a Ack)
	SendFinalDecision(fdchannel chan DecisionEnum, fd DecisionEnum)
}

// Worker represents a worker for the protocol.
type Worker interface {
	Start()
	DeliverCanCommit(c CanCommit)
	SendVote(v Vote)
	DeliverDecision(d Decision)
	SendACK(a Ack)
}

// Slice is a interface to represent a slice of messages.
type Slice interface {
	GetWorkID(idx int) WorkerID
}

// VoteSlice is the type to represent the slice of Vote.
type VoteSlice []Vote

// DecisionSlice is the type to represent the slice of Decision.
type DecisionSlice []Decision

// GetWorkID returns the workerID of each Vote in the slice of Vote.
func (votes VoteSlice) GetWorkID(idx int) WorkerID {
	return votes[idx].WorkerID
}

// GetWorkID returns the workerID of each Decision in the slice of Decision.
func (decisions DecisionSlice) GetWorkID(idx int) WorkerID {
	return decisions[idx].WorkerID
}

// Message is the interface to represent different types of messages.
type Message interface {
	GetWorkID() WorkerID
}

// GetWorkID returns the workerID of cancommit message.
func (cc CanCommit) GetWorkID() WorkerID {
	return cc.WorkerID
}

// GetWorkID returns the workerID of vote message.
func (v Vote) GetWorkID() WorkerID {
	return v.WorkerID
}

// GetVote returns the vote message value
func (v Vote) GetVote() VoteEnum {
	return v.VoteValue
}

// GetWorkID returns the workerID of decision message.
func (d Decision) GetWorkID() WorkerID {
	return d.WorkerID
}

// GetDecision returns the decision message value
func (d Decision) GetDecision() DecisionEnum {
	return d.DecisionValue
}

// search is a binary search function to find the message from a slice.
func Search(elems Slice, elem Message, low int, high int) int {
	if low > high {
		return -1
	} else {
		mid := (low + high - 1) / 2
		if elems.GetWorkID(mid) < elem.GetWorkID() {
			return Search(elems, elem, mid+1, high)
		} else if elems.GetWorkID(mid) > elem.GetWorkID() {
			return Search(elems, elem, low, mid-1)
		} else {
			return mid
		}
	}
}
