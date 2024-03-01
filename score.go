package main

import (
	"math"
	"sort"
	"time"
)

type Score struct {
	Committed []time.Time
	Score     float64
}

func LapseTime(created time.Time, committed time.Time, now time.Time) float64 {
	all := now.Sub(created).Seconds()
	com := now.Sub(committed).Seconds()
	return com / all
}

func CalcScore(ti float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -12.0*ti+12.0))
}

func (s *Score) Calc(now time.Time) {
	sort.Slice(s.Committed, func(i, j int) bool {
		return s.Committed[i].Before(s.Committed[j])
	})
	created := s.Committed[0]
	for _, c := range s.Committed {
		ti := LapseTime(created, c, now)
		s.Score += CalcScore(ti)
	}
}
