package main

import (
	"SDF/core"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	// Add Handlers
	_ "SDF/handlers"
)

func main() {
	// Load Environment Configuration
	dir, err := os.Getwd()
	core.CheckError(err)
	err = godotenv.Load(filepath.Join(dir, ".envrc"))
	core.CheckError(err)

  // core.RegisterStaticHandle("/assets/", "assets") Uncomment here if you have static files
	core.Serve("8000")
}

