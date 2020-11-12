package eat

import (
	// "encoding/json"
	"log"
	// "math/rand"
	// "strconv"
	// "time"
	"aces/plankton/utils"
)

type Mouth struct {
}

func NewMouth() *Mouth {
	return &Mouth{}
}

func (w *Mouth) ShapeInput() interface{} {
	return &utils.MouthInput{}
}

func (w *Mouth) Handler(input interface{}) (interface{}, error) {
	log.Println("mouth received input:", input.(*utils.MouthInput))
	return input, nil
}
