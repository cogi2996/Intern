package log

func Setup() {
	InitLogger(&Configuration{
		JSONFormat:      true,
		LogLevel:        "debug",
		StacktraceLevel: "fatal",
		Console:         &ConsoleConfiguration{},
	})
}
