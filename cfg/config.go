package cfg

import (
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
)

func SetupPath() string {
	var taskDir string
	scope := gap.NewScope(gap.User, "tasker")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatalf("Error getting data dirs: %v", err)
	}
	if len(dirs) == 0 {
		taskDir, _ = os.UserHomeDir()
	} else {
		taskDir = dirs[0]
	}
	err = initTaskerDir(taskDir)
	if err != nil {
		log.Fatalln("error initializing tasker directory:", err)
	}
	return taskDir
}

func initTaskerDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, 0o770)
	}
	return nil
}
