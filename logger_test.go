package logger

import (
	"testing"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}
func TestLogger(t *testing.T) {
	type Tweet struct {
		User    string
		Message string
	}

	Info("Tweet 0")
	Info("Tweet 1", Tweet{User: "User 1", Message: "Message 1"})
	Info("Tweet 2", Tweet{User: "User 2", Message: "Message 2"})
	Info("Tweet 3", Tweet{User: "User 3", Message: "Message 3"})
	Info("Tweet 4", Tweet{User: "User 4", Message: "Message 4"})
	Info("Tweet 5", Tweet{User: "User 5", Message: "Message 5"})
}
