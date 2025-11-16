package commands

func PointString(taskID string, currentTask string) string {
	if taskID == currentTask {
		return "-->"
	}
	return "   " // 3 spaces
}
