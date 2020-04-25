package singleton

import (
	"fmt"
	"sync"
)

var instance *singleton
var once sync.Once

type singleton struct{
	count int
}

func GetSingletonInstance() *singleton{
	once.Do(func(){
		instance = &singleton{}
	})

	return instance
}

func (s *singleton) PrintCount(){
	s.count+= 1
	fmt.Println(s.count)
}
