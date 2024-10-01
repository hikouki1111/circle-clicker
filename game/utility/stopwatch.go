package utility

import "time"

type Stopwatch struct {
	LastTime time.Time
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		time.Now(),
	}
}

func (s *Stopwatch) IsFinished(time int64, restart bool) bool {
	finished := s.Elapsed() >= time
	if finished && restart {
		s.Restart()
	}

	return finished
}

func (s *Stopwatch) Elapsed() int64 {
	return time.Now().UnixMilli() - s.LastTime.UnixMilli()
}

func (s *Stopwatch) Restart() {
	s.LastTime = time.Now()
}
