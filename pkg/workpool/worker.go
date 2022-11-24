package workpool

type worker struct {
	id       int
	work     chan Job
	workPool chan chan Job
	end      chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			w.workPool <- w.work
			select {
			case work := <-w.work:
				work.Execute()
			case <-w.end:
				return
			}
		}
	}()
}
func (w *worker) stop() {
	w.end <- struct{}{}
}

func newWorker(id int, workerQueue chan chan Job) *worker {
	work := &worker{
		id:       id,
		work:     make(chan Job),
		workPool: workerQueue,
		end:      make(chan struct{}, 1),
	}
	return work
}
