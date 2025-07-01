package worker;



type Task_Spawn struct {
}



func NewTask_Spawn() *Task_Spawn {
	return &Task_Spawn{};
}



func (task *Task_Spawn) Run() error {
print("RUN SPAWN TASK\n");
	return nil;
}
