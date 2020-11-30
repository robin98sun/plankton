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
	Size int    `json:"size,omitempty"`
}

type StomachInput struct {
	Cmd    string  `json:"cmd,omitempty"`
	Size   int     `json:"size,omitempty"`
	Pieces []int64 `json:"pieces,omitempty"`
}

func (w *Mouth) ShapeInput() interface{} {
	return &MouthInput{}
}

func (w *Mouth) Handler(inputInst interface{}) (interface{}, error) {
	var input *MouthInput
	input = inputInst.(*MouthInput)
	if input.Cmd == "gen and merge" && input.Size > 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		pieces := []int64{}
		for i := 0; i < input.Size; i++ {
			n := rand.Int63n(int64(input.Size * 100))
			pieces = append(pieces, n)
		}
		sort.Slice(pieces, func(i, j int) bool {
			return pieces[i] < pieces[j]
		})
		feedStomach := &StomachInput{
			Cmd:    input.Cmd,
			Size:   input.Size,
			Pieces: pieces,
		}
		return feedStomach, nil
	}
	return nil, nil
}
