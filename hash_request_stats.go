package main

// HashRequestStats is a struct containing stats relevent to hashing a password.
type HashRequestStats struct {
	HashRequests     int     `json:"total"`
	AverageTime      int     `json:"average"`
	totalTime        int
}

// Process stores password hashing statistics into the HashRequestStats object.
// This method is synchronized to prevent race conditions.
func (s *HashRequestStats) Process(time int) {
	mutex.Lock()
	s.HashRequests++
	s.totalTime += time
	s.AverageTime = s.totalTime / s.HashRequests
	mutex.Unlock()
}
