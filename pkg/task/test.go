package task

import (
	"fmt"

	"github.com/Uonx/falconscheduler/pkg/engine"
)

type Test struct {
	Name string
}

func (f *Test) Parse() (engine.ParseResult, error) {
	fmt.Printf("测试方法\n")
	var items []engine.Item
	items = append(items, engine.Item{
		Topic:   "test",
		Content: "测试队列添加数据",
	})
	return engine.ParseResult{Items: items}, nil
}
