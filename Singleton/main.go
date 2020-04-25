package main

import "github.com/sredxny/golang-design-patterns/Singleton/singleton"

func main(){
	mySingleton := singleton.GetSingletonInstance()
	mySingleton.PrintCount()

	anotherInstance := singleton.GetSingletonInstance()
	anotherInstance.PrintCount()

	thirdInstance := singleton.GetSingletonInstance()
	thirdInstance.PrintCount()
}
