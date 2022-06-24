package service

type ProcessorIf interface {
	Monitor()
}

type BaseProcessor struct{}

func (p BaseProcessor) Monitor() error {
	return nil
}
