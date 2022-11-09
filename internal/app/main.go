package main

func main() {
	err := setupRouter().Run(":8082")
	if err != nil {
		panic(err)
	}
}
