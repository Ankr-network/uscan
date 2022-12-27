package workpool

type Dispathcher interface {
	AddJob(Job)
	Stop()
}
