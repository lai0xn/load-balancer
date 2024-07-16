package main

func main() {
	b := NewBalancer(":8080")

	s, err := BuildServer("http://google.com:80")
	if err != nil {
		panic(err)
	}

	b.AddServer(s)

	i, err := BuildServer("http://instagram.com:80")
	if err != nil {
		panic(err)
	}

	b.AddServer(i)

	b.Listen()
}
