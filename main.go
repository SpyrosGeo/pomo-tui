package main

import (
	"os"
	"strconv"
	"time"
)


func main(){
	arguments:=os.Args
	if len(arguments) >2{
		println("Usage: go run main.go <minutes>")
		return
	}
	minutes,err:=strconv.Atoi(arguments[1])
	if(err!=nil){
		println("Error converting argument to integer:", err)
		return
	}
	if(minutes<=0){
		println("Please provide a positive integer for the timer.")
		return
	}
	startTimer(minutes)
}

func startTimer(minutes int) {
	if(minutes==1){
		println("Timer started for", minutes, "minute.")
	}else{	
	println("Timer started for", minutes, "minutes.")
	}
	timer := time.NewTimer(time.Duration(minutes) * time.Minute)
	<-timer.C
	println("Timer finished!")
}