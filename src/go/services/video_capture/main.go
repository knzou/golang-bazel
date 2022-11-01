package main

import (
	"fmt"
	"monorepo/src/go/services/video_capture/router"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine().Run(":8081")
}

func Check() {
	fmt.Println("Check")
}
