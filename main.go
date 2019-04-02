package main

func main() {
	ch := make(chan int, 1)
	analysisAvatarFiles(ch)
	<-ch
}
