package reporter

type Reporter interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	send(string)
}
