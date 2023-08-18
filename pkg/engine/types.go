package engine

type WorkerMethod interface {
	Parse() (ParseResult, error)
}

type ParseResult struct {
	Workers []WorkerMethod
	Items   []Item
}

type Item struct {
	Topic   string
	Content any
}
