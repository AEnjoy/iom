package global

type Task struct {
	Id         uint64 //0None 1software 2agent
	Type       uint64 //0host 1docker
	Args       string //args
	HasScripts uint32 //0None 1Python 2Shell 3yaml(container)
	ScriptsUrl string
}

var TaskList map[int]Task //clientID,Task

func TaskListInit() {
	TaskList = make(map[int]Task)
}
