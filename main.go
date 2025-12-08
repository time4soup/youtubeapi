package main

func main() {
	//add: handle error from repl()
	err := repl()
	if err != nil {
		panic(err)
	}
}
