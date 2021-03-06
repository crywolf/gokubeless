package kubeless

import (
  "fmt"

  "github.com/kubeless/kubeless/pkg/functions"
)

// Handler sample function with data
func Handler(event functions.Event, context functions.Context) (string, error) {
  fmt.Println(event)
  return "Hello, I heard you said: " + event.Data, nil
}