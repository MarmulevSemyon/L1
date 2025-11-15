package main

var justString string

func main() {
	v := string(1 << 10)
	justString = v[:100]
}
