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
	Size int64  `json:"size,omitempty"`
	StartPoint int64 `json:"startPoint,omitempty"`
	EndPoint   int64 `json:"endPoint,omitempty"`
	MaxEatTime int64 `json:"maxEatTime,omitempty"`
	MinEatTime int64 `json:"minEatTime,omitempty"`
	DigestTime float64 `json:"digestTime,omitempty"`
	DigestFactor int64 `json:"digestFactor,omitempty"`
}

type StomachInput struct {
	Cmd    string   `json:"cmd,omitempty"`
	EatTime float64 `json:"eatTime,omitempty"`
	Size    int64   `json:"size,omitempty"`
	Pieces  []int64 `json:"pieces,omitempty"`
	DigestTime float64 `json:"digestTime,omitempty"`
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
		EatTime:   float64(input.Size),
	}
	// do some job
	if (input.Cmd == "find prime numbers and count") {

		is_prime_number := func(n int64) bool {
			for i := int64(2); i<n; i++ {
				if n % i == 0 {
					return false
				}
			}
			return true
		}

		startTime := time.Now()

		prime_number_list := []int64{}
		for i := input.StartPoint; i<= input.EndPoint; i++ {
			if is_prime_number(i) {
				prime_number_list = append(prime_number_list, i)
			}
		}

		endTime := time.Now()
		// feedStomach.Pieces = prime_number_list
		// feedStomach.Size   = input.EndPoint - input.StartPoint + 1
		feedStomach.Size = int64(len(prime_number_list))
		feedStomach.EatTime = float64((endTime.Sub(startTime))*time.Millisecond)
		feedStomach.Cmd = "count"

	} else if (input.Cmd == "gen and merge" || input.Cmd == "gen and merge and wait") && input.Size > 0 {
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
		feedStomach.EatTime = float64(n)
	}

	// forward digest options
	feedStomach.DigestTime = input.DigestTime
	feedStomach.DigestFactor = input.DigestFactor

	// done
	return feedStomach, nil
}
