package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	// Background: root context, no deadline, no cancellation
	background := context.Background()
	fmt.Println(background)

	// TODO: placeholder context, use when unsure which context to use
	todo := context.TODO()
	fmt.Println(todo)
}