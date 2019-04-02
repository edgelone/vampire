package main

import (
	"sync"
)

func main() {

	var wg sync.WaitGroup

	analysisAvatarFiles(wg)

	wg.Wait()

}
