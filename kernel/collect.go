package kernel

import (
	"aces/video-comb-aggregator/utils"
	"encoding/json"
	"fmt"
	"log"
)

type Core struct {
	Conf     *utils.Conf
	status   string
	subtasks map[string][]*Result
}

func NewCore(conf *utils.Conf) *Core {
	c := &Core{
		Conf:     conf,
		status:   "init",
		subtasks: make(map[string][]*Result),
	}
	for _, subtaskID := range conf.Subtasks {
		c.subtasks[subtaskID] = nil
	}
	return c
}

type Result struct {
	Image      string
	Likelihood float32
}

type ResultSet struct {
	TaskID    string
	SubtaskID string
	Results   []*Result
}

func (c *Core) finished() int {
	if c == nil || c.subtasks == nil || len(c.subtasks) == 0 {
		return 0
	}
	finished := 0
	for _, results := range c.subtasks {
		if results != nil {
			finished++
		}
	}
	return finished
}
func (c *Core) Status() string {
	if c == nil || c.subtasks == nil {
		return "not initiated"
	} else if len(c.subtasks) == 0 {
		return "no subtask"
	} else {
		finished := c.finished()
		return fmt.Sprintf("total %v subtasks, %v finished, %v still waiting", len(c.subtasks), finished, len(c.subtasks)-finished)
	}
}

func (c *Core) ReceiveResults(resultBytes []byte) {
	c.status = "Receiving"
	resultset := &ResultSet{}
	json.Unmarshal(resultBytes, resultset)
	resultstr, _ := json.MarshalIndent(resultset, "", "  ")
	log.Println("received sub task result:", string(resultstr))
	c.subtasks[resultset.SubtaskID] = resultset.Results
	if c.finished() == len(c.subtasks) {
		c.reportToMaster()
		c.reportToReducer(nil)
		log.Println("done")
		c.status = "DONE"
	}
}

func (c *Core) reportToMaster() {
	if !c.Conf.MasterNode.IsValid() {
		return
	}
	log.Println("reporting to master")
	c.status = "Reporting to Master"
}

func (c *Core) reportToReducer(results []*Result) {
	if !c.Conf.AggregatorNode.IsValid() {
		return
	}
	log.Println("reporting to reducer, results:", results)
	c.status = "Reporting to Reducer"
}
