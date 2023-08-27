package examples

import (
	"testing"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func TestEx01(t *testing.T) {
	type testrt struct {
		ID      int
		Content string
		GID     int
	}

	cl := manticore.NewClient()
	cl.SetServer("127.0.0.1", 9312)
	ok, err := cl.Open()
	if err != nil {
		t.Fatalf("open connection: %v", err)
	}

	t.Logf("OK: %v", ok)

	results, err := cl.Sphinxql("SELECT * FROM testrt")
	if err != nil {
		t.Fatalf("Sphinxql select: %v", err)
	}

	// Можно получить ответ, как если делать sql запрос в cli утилите
	//for _, r := range results {
	//	t.Logf("Res: %v", r.String())
	//}

	for _, r := range results {
		for _, e := range r.Rows {
			for c := range e {
				t.Logf("c: %v", c)
			}

		}
	}

}
