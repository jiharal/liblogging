package liblogging

import (
	"testing"
)

func TestLiblogging(t *testing.T) {
	newLog := New()
	newLog.Out.Println("Hi")
	newLog.Err.Println("Hi")
	newLog.Field.Println("Hi")
}