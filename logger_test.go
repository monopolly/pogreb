package pogreb

//testing
import (
	"fmt"
	"strings"
	"testing"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestMain(u *testing.T) {
	__(u)

	/* New("temp1", func(s string) {
		fmt.Print(s)
	})

	b, _ := New("temp2")
	//go Closer(a, b)

	for i := 0; i < 100000; i++ {
		b.Add(fmt.Sprint(i), []byte("ok"))
	}

	fmt.Println("done")
	select {} */
}

func BenchmarkPrint(u *testing.B) {
	u.ReportAllocs()
	u.ResetTimer()
	for n := 0; n < u.N; n++ {

	}
}

func BenchmarkNoPrint(u *testing.B) {
	u.ReportAllocs()
	u.ResetTimer()
	for n := 0; n < u.N; n++ {

	}
}

func __(u *testing.T) {
	fmt.Printf("\033[1;32m%s\033[0m\n", strings.ReplaceAll(u.Name(), "Test", ""))
}
