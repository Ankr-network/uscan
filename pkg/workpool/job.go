package workpool

type Job interface {
	Execute()
}
