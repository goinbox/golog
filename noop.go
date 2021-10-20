package golog

type NoopLogger struct {
}

func (l *NoopLogger) Debug(msg string, fields ...*Field) {
}

func (l *NoopLogger) Info(msg string, fields ...*Field) {
}

func (l *NoopLogger) Notice(msg string, fields ...*Field) {
}

func (l *NoopLogger) Warning(msg string, fields ...*Field) {
}

func (l *NoopLogger) Error(msg string, fields ...*Field) {
}

func (l *NoopLogger) Critical(msg string, fields ...*Field) {
}

func (l *NoopLogger) Alert(msg string, fields ...*Field) {
}

func (l *NoopLogger) Emergency(msg string, fields ...*Field) {
}

func (l *NoopLogger) With(fields ...*Field) Logger {
	return l
}

type NoopWriter struct {
}

func (n *NoopWriter) Write(msg []byte) (int, error) {
	return 0, nil
}

func (n *NoopWriter) Flush() error {
	return nil
}

func (n *NoopWriter) Free() {
}
