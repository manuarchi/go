package hello
import (
    "testing"
	"fmt"
)
 
func TestHello(t *testing.T) {
	s := Hello()
	if s=="hello world" {
		fmt.Println(s)
	 }else {
	 	t.Error(s)
	}
}