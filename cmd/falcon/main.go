package main

import (
	"github.com/Uonx/falconscheduler/pkg/engine"
	"github.com/Uonx/falconscheduler/pkg/persist"
	"github.com/Uonx/falconscheduler/pkg/scheduler"
	"github.com/Uonx/falconscheduler/pkg/task"
)

func main() {
	itemSaver, _ := persist.ItemSaver()

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
	}
	e.Run(&task.Test{})
}
