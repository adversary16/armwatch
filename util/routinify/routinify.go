package routinify

import (
	"fmt"
	"time"
)

func Add(stack []func() (int, error), updateCycleMsec int) {
	updateCycle := time.Duration(updateCycleMsec * int(time.Millisecond))
	cycleTicker := time.NewTicker(updateCycle)
	for _, i := range stack {
		for {
			select {
			case <-cycleTicker.C:
				go func() {
					_, err := i()
					if err != nil {
						fmt.Println(err)
					}
				}()
			}
		}
	}
}
