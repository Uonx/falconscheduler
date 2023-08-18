package persist

import (
	"fmt"

	"github.com/Uonx/falconscheduler/pkg/engine"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			fmt.Printf("item topic:%s content:%v \n", item.Topic, item.Content)
		}
	}()
	return out, nil
}
