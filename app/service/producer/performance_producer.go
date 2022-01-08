package producer

var (
	PerformanceProducer chan<- []byte
	TcpServerProducer   chan<- []byte
)
