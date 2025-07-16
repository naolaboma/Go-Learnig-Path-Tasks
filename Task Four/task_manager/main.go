package main

import (
	"fmt"
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	fmt.Println("Task manager apo running on http://localhost:8080")
	r.Run(":8080")
}
