package main

//We can use backtick(`) to create a multiling string. But be careful - any spacing you use in the string to retain indentation will also be present in the final string.
import "fmt"

func main() {
	str := `This is a
	multiling
	string.`
	fmt.Println(str)
}
