package progress_test

import (
	"fmt"
	"os"

	"github.com/drylikov/go_progress"
)

func Example() {
	b := progress.NewInt(50)

	for i := 0; i <= 50; i++ {
		b.ValueInt(i)
		b.Text(fmt.Sprintf("iteration %d", i))
		b.WriteTo(os.Stdout)
	}
}
