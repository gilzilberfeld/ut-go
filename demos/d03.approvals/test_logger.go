package demos

import (
	"strings"
)

type TestLogger struct {
    sb strings.Builder
}

func (t *TestLogger) Append(key string, display string) {
    t.sb.WriteString("Pressed " + key + ", Display shows: " + display + "\n")
}

func (t *TestLogger) GetAll() string {
    return t.sb.String()
}

