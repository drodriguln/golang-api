package model

import "sync"

var mutex = &sync.Mutex{}
var Stats = HashRequestStats{}

// HashRequestStats is a struct containing stats relevent to hashing a password.
type HashRequestStats struct {
	HashRequests     int     `json:"total"`
	AverageTime      int     `json:"average"`
	TotalTime        int     `json:"-"`
}

// Process stores password hashing statistics into the HashRequestStats object.
// This method is synchronized to prevent race conditions.
func (s *HashRequestStats) Process(time int) {
	mutex.Lock()
	s.HashRequests++
	s.TotalTime += time
	s.AverageTime = s.TotalTime / s.HashRequests
	mutex.Unlock()
}
