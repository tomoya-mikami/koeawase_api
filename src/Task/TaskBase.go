package Task

type TaskInterface interface {
	Execute(taskName string, args []string)
}
