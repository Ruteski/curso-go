package main

// go env
// https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
// go tool dist list
// go build -o <nome que eu quero>
// entendimento gcc
func main() {
	println("Compilando no go")
	println("GOOS=windows go build main.go")
	println("GOOS=linux go build main.go")
	println("GOOS=darwin go build main.go")               // para mac
	println("GOARCH=amd64 GOOS=windows go build main.go") // para windows com amd
	println("quando tiver um modulo -> go build")
	println("quando tiver um modulo e querer mudar o nome -> go build -o nome")
}
