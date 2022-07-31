package pogreb

//testing
import (
	"fmt"
	"os"
	"strings"
	"testing"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestMain(u *testing.T) {
	__(u)
	defer os.RemoveAll("temp1")
	d, _ := New("temp1", func(s string) {
		fmt.Print(s)
	})
	d.Close()

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
