package main

//Go 1.10 has a new strings.Builder! It is also an efficient way
//While Go does allow you to concatenate strings with the + operator, this can become pretty inefficient when concatenating a lot of strings together. It is much more efficient to use a bytes.Buffer and then convert it to a string once you have concatenated everything.
import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(randString())
	}
	fmt.Println(b.String())
}

func randString() string {
	return "abc-123-ABC-"
}
