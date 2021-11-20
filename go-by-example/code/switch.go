package gobyex

import (
	"fmt"
	"time"
)

func Switch() {

	day := time.Now().Weekday()
	fmt.Println(day)

	switch day {
	case time.Friday:
		fmt.Println("yay weekend")
	case time.Monday:
		fmt.Println("work")
	default:
		fmt.Println("waht day is this?")
	}

	//if else using swtich

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	default:
		fmt.Println("evening")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("no idea")
		}
	}

	whatAmI(false)
	whatAmI("string")
}
