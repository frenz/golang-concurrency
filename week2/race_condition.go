package main

import (
	"time"
	"fmt"
)


func goroutine(s string, i *int) {
	*i = *i + 1
	fmt.Println(s, *i)	
}


func main() {
	/*
	possible outputs:

		Assignment week4: race_condition.go
		first 1
		second 2
		end

		Assignment week4: race_condition.go
		second 2
		first 2
		end
		
		Assignment week4: race_condition.go
		second 2
		first 1
		end
	there is a race condition to the access to i
	*/
	fmt.Println("Assignment week4: race_condition.go")
	i := 0
	go goroutine("first", &i)
	go goroutine("second", &i)
	time.Sleep(time.Millisecond * 100) 
	fmt.Println("end")
}
