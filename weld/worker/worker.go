package worker;



type Worker struct {
}



func New() *Worker {
	return &Worker{};
}

func (work *Worker) Run() error {
	return nil;
}
