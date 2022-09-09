package main

import (
	"encoding/json"
	"fmt"
	"github.com/panjf2000/ants"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type SvcConfig struct {
	MaxPool       int    `json:"max_pool"`
	MinPool       int    `json:"min_pool"`
	PoolIncrement int    `json:"pool_increment"`
	AutoTunePool  string `json:"auto_tune_pool"`
}

var runningTasks int32

func main() {
	var wg sync.WaitGroup
	// number of tasks will be run
	numberOfTasks := 100

	// get config from config file
	config, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	var svcConfig SvcConfig
	jsonParser := json.NewDecoder(config)
	if err := jsonParser.Decode(&svcConfig); err != nil {
		panic(err)
	}

	// pool initialization
	pool, _ := ants.NewPool(svcConfig.MinPool)

	// auto tune ygy
	go AutoTunePool(pool, svcConfig, &wg)

	// do task
	for i := 0; i < numberOfTasks; i++ {
		pool.Submit(
			func() {
				DoHeavyTask(&wg, i)
			},
		)

		fmt.Println("Running tasks: ", atomic.LoadInt32(&runningTasks))
	}

	wg.Wait()

}

// auto configure pool every ... time
func AutoTunePool(pool *ants.Pool, svcConfig SvcConfig, wg *sync.WaitGroup) {
	duration, _ := time.ParseDuration(svcConfig.AutoTunePool)
	wg.Add(1)
	for {
		if int(atomic.LoadInt32(&runningTasks)) >= (pool.Cap() - 1) {
			// tune pool
			if pool.Cap()+svcConfig.PoolIncrement <= svcConfig.MaxPool {

				newPool := pool.Cap() + svcConfig.PoolIncrement
				pool.Tune(uint(newPool))
				fmt.Println("Pool cap configured to ", pool.Cap())
			}
		} else if pool.Cap()-svcConfig.PoolIncrement >= svcConfig.MinPool {

			newPool := pool.Cap() - svcConfig.PoolIncrement
			pool.Tune(uint(newPool))
			fmt.Println("Pool cap configured to ", pool.Cap())

		}
		time.Sleep(duration)
	}
}

func DoHeavyTask(wg *sync.WaitGroup, arg int) {
	atomic.AddInt32(&runningTasks, 1)
	wg.Add(1)
	//fmt.Println("Task ", arg, " started")
	time.Sleep(1 * time.Second)
	//fmt.Println("Task ", arg, " finished")
	atomic.AddInt32(&runningTasks, -1)
	wg.Done()
}
