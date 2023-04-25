package digest

import (
	"log"
	"time"
)

type Stomach struct {
}

func NewStomach() *Stomach {
	return &Stomach{}
}

type StomachInput struct {
	Cmd    string  `json:"cmd,omitempty"`
	EatTime   float64  `json:"eatTime,omitempty"`
	Size      int64  `json:"size,omitempty"`
	Pieces []int64 `json:"pieces,omitempty"`
	DigestTime float64`json:"digestTime,omitempty"`
	DigestFactor int64 `json:"digestFactor,omitempty"`
}

func (w *Stomach) ShapeResultOfSubtask() interface{} {
	return &StomachInput{}
}

func (w *Stomach) ShapeCumulation() interface{} {
	return &StomachInput{}
}

func (w *Stomach) Handler(cumulationInst interface{}, previousResults []interface{}, subtaskResultInst interface{}) (interface{}, error) {
	if subtaskResultInst == nil {
		return cumulationInst, nil
	}
	subtaskResult := subtaskResultInst.(*StomachInput)
	result := &StomachInput{
		Cmd: subtaskResult.Cmd,
		EatTime: subtaskResult.EatTime,
	}
	var cumulation *StomachInput
	if cumulationInst != nil {
		cumulation = cumulationInst.(*StomachInput)
	}

	// aggregate subtasks
	if subtaskResult.Cmd == "count" {
		if cumulation == nil {
			result = subtaskResult
		} else {
			result.Size = cumulation.Size + subtaskResult.Size
			log.Println("[stomach] stomach counted the subtask result")
		}	
	} else if subtaskResult.Cmd == "merge" || subtaskResult.Cmd == "gen and merge" || subtaskResult.Cmd == "gen and merge and wait"{
		if cumulation == nil {
			result = subtaskResult
		} else {
			i := 0
			j := 0

			log.Println("[stomach] start merging")
			log.Printf("[stomach] length of cumulation: %v, length of subtask: %v", len(cumulation.Pieces), len(subtaskResult.Pieces))
			for i < len(cumulation.Pieces) && j < len(subtaskResult.Pieces) {
				if subtaskResult.Pieces[j] < cumulation.Pieces[i] {
					result.Pieces = append(result.Pieces, subtaskResult.Pieces[j])
					j++
				} else {
					result.Pieces = append(result.Pieces, cumulation.Pieces[i])
					i++
				}
			}
			log.Println("[stomach] common part done")
			for ; i < len(cumulation.Pieces); i++ {
				result.Pieces = append(result.Pieces, cumulation.Pieces[i])
			}
			log.Println("[stomach] cleared cumulation")
			for ; j < len(subtaskResult.Pieces); j++ {
				result.Pieces = append(result.Pieces, subtaskResult.Pieces[j])
			}
			result.Size = subtaskResult.Size + cumulation.Size
			log.Println("[stomach] cleared subtask data")
			log.Println("[stomach] merging done")
		}
	}

	// wait some time
	if subtaskResult.Cmd == "service time" || subtaskResult.Cmd == "gen and merge and wait" {
		if subtaskResult.DigestTime > 0 {
			timeToWait := subtaskResult.DigestTime
			if subtaskResult.DigestFactor > 0 {
				previousSubtaskCount := float64(len(previousResults))
				timeToWait += previousSubtaskCount * float64(subtaskResult.DigestFactor) * timeToWait
				result.DigestFactor = subtaskResult.DigestFactor
			}
			time.Sleep(time.Duration(timeToWait) * time.Millisecond)
			result.DigestTime = timeToWait
		}
	}

	return result, nil
}
