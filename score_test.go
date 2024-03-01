package main

import (
	"testing"
	"time"
)

func TestLapseTime(t *testing.T) {
	now := time.Now()
	t1 := now.Add(-time.Hour)
	t0 := t1.Add(-time.Hour)
	result := LapseTime(t0, t1, now)
	if result != 0.5 {
		t.Errorf("expected: 0.5, but %v", result)
	}

	result = LapseTime(t0, t0, now)
	if result != 0.0 {
		t.Errorf("expected: 0.0, but %v", result)
	}

	result = LapseTime(t0, now, now)
	result = LapseTime(t0, t0, now)
	if result != 1.0 {
		t.Errorf("expected: 1.0, but %v", result)
	}
}

func TestCalcScore(t *testing.T) {
	result := CalcScore(0.0)
	if result > 7.0e-6 {
		t.Errorf("expected: less than 7.0E-6, but %v", result)
	}
	result = CalcScore(1.0)
	if result != 0.5 {
		t.Errorf("expected: 0.5, but %v", result)
	}
}
