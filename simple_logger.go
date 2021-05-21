/**
* @file logger.go
* @author ligang
* @date 2016-02-04
 */

package golog

type simpleLogger struct {
	writer   Writer
	formater Formater

	glevel int
}

func NewSimpleLogger(writer Writer, formater Formater) *simpleLogger {
	return &simpleLogger{
		writer:   writer,
		formater: formater,

		glevel: LevelInfo,
	}
}

func (s *simpleLogger) SetLogLevel(level int) *simpleLogger {
	_, ok := LogLevels[level]
	if ok {
		s.glevel = level
	}

	return s
}

func (s *simpleLogger) Debug(msg []byte) {
	_ = s.Log(LevelDebug, msg)
}

func (s *simpleLogger) Info(msg []byte) {
	_ = s.Log(LevelInfo, msg)
}

func (s *simpleLogger) Notice(msg []byte) {
	_ = s.Log(LevelNotice, msg)
}

func (s *simpleLogger) Warning(msg []byte) {
	_ = s.Log(LevelWarning, msg)
}

func (s *simpleLogger) Error(msg []byte) {
	_ = s.Log(LevelError, msg)
}

func (s *simpleLogger) Critical(msg []byte) {
	_ = s.Log(LevelCritical, msg)
}

func (s *simpleLogger) Alert(msg []byte) {
	_ = s.Log(LevelAlert, msg)
}

func (s *simpleLogger) Emergency(msg []byte) {
	_ = s.Log(LevelEmergency, msg)
}

func (s *simpleLogger) Log(level int, msg []byte) error {
	if level > s.glevel {
		return nil
	}

	_, err := s.writer.Write(s.formater.Format(level, append(msg, '\n')))

	return err
}
