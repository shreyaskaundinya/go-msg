package queue

import "sync"

type Queue struct {
	Data []string
	Top  int
	Lock sync.RWMutex
}