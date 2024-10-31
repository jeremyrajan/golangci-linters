package testpackage

// Simulating logger interfaces without external dependencies
type Logger struct{}

func (l *Logger) Info(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field) {}
func (l *Logger) Debug(msg string, fields ...Field) {}

type Field struct {
	Key   string
	Value interface{}
}

func String(key string, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func ExampleUsage() {
	log := &Logger{}

	// These should trigger warnings
	log.Info("just a message")   // want "logging calls must include metadata fields"
	log.Error("another message") // want "logging calls must include metadata fields"

	// These should pass
	log.Info("user logged in", String("user_id", "123"))
	log.Error("operation failed",
		String("error", "connection timeout"),
		Int("retry_count", 3),
	)
}
