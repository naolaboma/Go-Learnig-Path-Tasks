package main

import (
	"library_management/controllers"
)

func main() {
	libraryCtrl := controllers.NewLibraryController()
	libraryCtrl.Run()
}
