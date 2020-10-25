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

func (w *Stomach) NewInput() interface{} {
	return &utils.StomachInput{}
}

func (w *Stomach) Handler(input interface{}) (interface{}, error) {
	log.Println("stomach received input:", input.(*utils.StomachInput))
	return input, nil
}
