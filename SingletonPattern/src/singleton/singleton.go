package singleton

import "sync"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// Lazy initialization
var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{count: 100}
	})
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

// Eager Initialization:
//func init() {
//	instance = &singleton{count: 100}
//}
