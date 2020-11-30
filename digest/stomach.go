package digest

import (
	"log"
)

type Stomach struct {
}

func NewStomach() *Stomach {
	return &Stomach{}
}

type StomachInput struct {
	Cmd    string  `json:"cmd,omitempty"`
	Size   int     `json:"size,omitempty"`
	Pieces []int64 `json:"pieces,omitempty"`
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
	var result *StomachInput
	var cumulation *StomachInput
	if cumulationInst != nil {
		cumulation = cumulationInst.(*StomachInput)
	}
	if subtaskResult.Cmd == "gen and merge" {
		if cumulation == nil {
			result = subtaskResult
		} else {
			i := 0
			j := 0

			result = &StomachInput{
				Cmd: subtaskResult.Cmd,
			}
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
			result.Size += subtaskResult.Size
			log.Println("[stomach] cleared subtask data")
			log.Println("[stomach] merging done")
		}
	}

	return result, nil
}
