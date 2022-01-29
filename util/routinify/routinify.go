package routinify

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Add(stack []func() (int, error), wg *sync.WaitGroup, updateCycleMsec int) {
	loopPeriod, err := time.ParseDuration(strings.Join([]string{strconv.Itoa(updateCycleMsec), "ms"}, ""))
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range stack {
		wg.Add(1)
		go func() {
			for {
				_, err := i()
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(loopPeriod)
			}
		}()
	}
	// wg.Add(1)
}
