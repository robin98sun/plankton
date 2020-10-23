package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Addr     string
	Port     int
	Protocol string
}

func (n *Node) Key() string {
	return fmt.Sprintf("%v:%v", n.Addr, n.Port)
}

func (n *Node) IsValid() bool {
	return n != nil && n.Addr != "" && n.Port > 0
}

type VideoStream struct {
	URL string
}

type Conf struct {
	TaskID         string
	SubtaskID      string
	SelfNode       *Node
	AggregatorNode *Node
	MasterNode     *Node // to report the completion of sub-task
	TTL            int64 // in milliseconds, if job can not done in TTL, the pod would be forced to clean
	Subtasks       []string
}

// ReadConfFromEnv Read configuration from environment variables
func ReadConfFromEnv() *Conf {
	allEnv := os.Environ()
	log.Println("all environment variables:\n", allEnv)
	conf := &Conf{
		SelfNode:       &Node{},
		AggregatorNode: &Node{},
		MasterNode:     &Node{},
		Subtasks:       []string{},
	}
	conf.TaskID = os.Getenv("JADE_TASKID")
	conf.SubtaskID = os.Getenv("JADE_SUBTASKID")
	conf.SelfNode.Protocol = os.Getenv("JADE_SELFNODE_PROTOCOL")
	conf.SelfNode.Addr = os.Getenv("JADE_SELFNODE_ADDR")
	conf.SelfNode.Port, _ = strconv.Atoi(os.Getenv("JADE_SELFNODE_PORT"))
	conf.MasterNode.Protocol = os.Getenv("JADE_MASTERNODE_PROTOCOL")
	conf.MasterNode.Addr = os.Getenv("JADE_MASTERNODE_ADDR")
	conf.MasterNode.Port, _ = strconv.Atoi(os.Getenv("JADE_MASTERNODE_PORT"))
	conf.AggregatorNode.Protocol = os.Getenv("JADE_AGGREGATORNODE_PROTOCOL")
	conf.AggregatorNode.Addr = os.Getenv("JADE_AGGREGATORNODE_ADDR")
	conf.AggregatorNode.Port, _ = strconv.Atoi(os.Getenv("JADE_AGGREGATORNODE_PORT"))
	conf.TTL, _ = strconv.ParseInt(os.Getenv("JADE_TTL"), 10, 64)
	subtaskstr := os.Getenv("JADE_SUBTASKS")
	conf.Subtasks = strings.Split(subtaskstr, ",")
	return conf
}
