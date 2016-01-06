package main

import "github.com/xlab/portaudio-go/portaudio"

func paError(err portaudio.Error) bool {
	return portaudio.ErrorCode(err) != portaudio.PaNoError
}

type uint32Row struct {
	N     uint32
	Value uint32
}

type uint32Rows []uint32Row

// SearchFirst is used to get number of audio packets per chunk N
// provided the following list of pairs:
//
// #0 : 5 packets per chunk starting @chunk #1
// #1 : 2 packets per chunk starting @chunk #131
func (s uint32Rows) SearchFirst(N uint32) (Value uint32) {
	if len(s) == 0 {
		return 0
	}
	for i := range s {
		if s[i].N < N {
			continue
		} else if s[i].N == N {
			return s[i].Value
		}
		return s[i-1].Value
	}
	return s[len(s)-1].Value
}

// SearchLast is used to get the duration of audio frame N
// provided the following list of pairs:
//
// #0 : 651 frames with duration 4096 units
// #1 : 1 frames with duration 672 units
func (s uint32Rows) SearchLast(N uint32) (Value uint32) {
	if len(s) == 0 {
		return 0
	}
	for i := range s {
		if n := s[i].N; N < n {
			return s[i].Value
		} else if N = N - n; N >= 0 {
			continue
		}
	}
	return s[len(s)-1].Value
}
