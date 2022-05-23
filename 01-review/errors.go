package main

import (
	"errors"
	"fmt"
	"time"
)

/*
Errors son interfaces
*/

func terror()error{
	if time.Now().Unix() %2 ==0{
		return nil
	}
	return &MyError{Msg: "this is bad"}
}

type MyError struct{
	Msg string
}
func (e *MyError) Error() string{
	return e.Msg
}

func main(){
	err := terror()

	if err != nil{
		panic("time is not even")
	}

	res, err := division(10,2)
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	
}

func division(a,b int)(float32, error){
	if b==0{
		return 0.0, errors.New("Cannot divide by zero")
	}

	return float32(a)/float32(b), nil
}