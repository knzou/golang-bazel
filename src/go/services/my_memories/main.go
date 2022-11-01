package main

import (
	"fmt"
	"monorepo/src/go/services/my_memories/router"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine().Run(":8080")
}

func Check() {
	fmt.Println("Check")
}
