package worker;

import(
);



type Worker struct {
}



func New() *Worker {
	return &Worker{};
}

func (worker *Worker) Run() error {
	return nil;
}
