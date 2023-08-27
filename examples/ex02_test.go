package examples

import (
	"testing"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func TestEx02(t *testing.T) {
	cl := manticore.NewClient()
	cl.SetServer("127.0.0.1", 9312)
	_, err := cl.Open()
	if err != nil {
		t.Fatalf("open connection: %v", err)
	}

	// Создаем запрос
	q := manticore.NewSearch("content", "testrt", "")
	result, err := cl.RunQuery(q)
	if err != nil {
		t.Fatalf("run search query: %v", err)
	}

	t.Logf("Res: %v", result)

	// Можно добавить условие и повторить запрос
	q.AddFilterExpression("gid > 10 AND gid < 20", false)
	result2, err2 := cl.RunQuery(q)
	if err2 != nil {
		t.Fatalf("run query-2: %v", err2)
	}

	t.Logf("Result-2: %v", result2)

	// Не знаю, как использовать эту информацию.
	//for _, f := range result2.Fields {
	//	t.Logf("Field: %v", f)
	//}

	// Не знаю, как использовать эту информацию.
	//for _, r := range result2.Matches {
	//	t.Logf("Matche: %v", r)
	//}

	//t.Logf("Res: %s", result2.String())
}
