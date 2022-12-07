package day6

import (
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/utils"
)

func detectMarker(data string, windowSize int) (string, int) {
	chars := collections.NewHashSetOfNum[byte]()

	for i := windowSize - 1; i < len(data); i++ {
		window := data[(i - (windowSize - 1)):(i + 1)]

		utils.AddSliceTo[byte]([]byte(window), chars)

		if chars.Size() == windowSize {
			return window, i + 1
		}
		chars.Clear()
	}
	panic("data exhausted without finding the marker!")
}
