package digest

import (
	// "encoding/json"
	"log"
	// "math/rand"
	// "strconv"
	// "time"
	"aces/plankton/utils"
)

type Stomach struct {
}

func NewStomach() *Stomach {
	return &Stomach{}
}

func (w *Stomach) ShapeResultOfSubtask() interface{} {
	return &utils.StomachInput{}
}

func (w *Stomach) ShapeCumulation() interface{} {
	return &utils.StomachInput{}
}

func (w *Stomach) Handler(cumulation interface{}, previousResults []interface{}, subtaskResult interface{}) (interface{}, error) {
	log.Println("stomach received input:", subtaskResult.(*utils.StomachInput))
	var result *utils.StomachInput
	result = cumulation.(*utils.StomachInput)
	if cumulation == nil {
		result = subtaskResult.(*utils.StomachInput)
	} else if subtaskResult != nil {
		result.Key1 += ", " + subtaskResult.(*utils.StomachInput).Key1
		result.Key3 += ", " + subtaskResult.(*utils.StomachInput).Key3
		result.Key2 += subtaskResult.(*utils.StomachInput).Key2
	}

	return result, nil
}
