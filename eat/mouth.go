package eat

import (
	"math/rand"
	"sort"
	"time"
)

type Mouth struct {
}

func NewMouth() *Mouth {
	return &Mouth{}
}

type MouthInput struct {
	Cmd  string `json:"cmd,omitempty"`
	Size int64    `json:"size,omitempty"`
	MaxEatTime  int64    `json:"maxEatTime,omitempty"`
	MinEatTime  int64    `json:"minEatTime,omitempty"`
	DigestTime int64 `json:"digestTime,omitempty"`
	DigestFactor int64 `json:"digestFactor,omitempty"`
}

type StomachInput struct {
	Cmd    string  `json:"cmd,omitempty"`
	EatTime   int64  `json:"eatTime,omitempty"`
	Size      int64  `json:"size,omitempty"`
	Pieces []int64 `json:"pieces,omitempty"`
	DigestTime int64 `json:"digestTime,omitempty"`
	DigestFactor int64 `json:"digestFactor,omitempty"`
}

func (w *Mouth) ShapeInput() interface{} {
	return &MouthInput{}
}

func (w *Mouth) Handler(inputInst interface{}) (interface{}, error) {
	var input *MouthInput
	input = inputInst.(*MouthInput)

	feedStomach := &StomachInput{
		Cmd:    input.Cmd,
		EatTime:   input.Size,
	}
	// do some job
	if (input.Cmd == "gen and merge" || input.Cmd == "gen and merge and wait") && input.Size > 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		pieces := []int64{}
		for i := int64(0); i < input.Size; i++ {
			n := rand.Int63n(input.Size * 100)
			pieces = append(pieces, n)
		}
		sort.Slice(pieces, func(i, j int) bool {
			return pieces[i] < pieces[j]
		})
		feedStomach.Pieces = pieces
		feedStomach.EatTime = 0
	} 

	// wait some time
	if input.Cmd == "service time" || input.Cmd == "gen and merge and wait" {
		rand.Seed(time.Now().UTC().UnixNano())
		n := int64(0)
		if input.MaxEatTime >= input.MinEatTime && input.MinEatTime > 0 {
			if input.MinEatTime == input.MaxEatTime {
				n = input.MinEatTime
			} else {
				n = rand.Int63n(input.MaxEatTime-input.MinEatTime)
				n += int64(input.MinEatTime)
			}
		}
		time.Sleep(time.Duration(n) * time.Millisecond)
		feedStomach.EatTime = n	
	}

	// forward digest options
	feedStomach.DigestTime = input.DigestTime
	feedStomach.DigestFactor = input.DigestFactor

	// done
	return feedStomach, nil
}
