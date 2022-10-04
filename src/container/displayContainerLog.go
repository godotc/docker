package container

import (
	"docker/src/utils_"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func DisplayContainerLog(containerName string) {
	logPath := filepath.Join(ROOT_FOLDER_PATH_PREFIX, containerName, LOG_FILENAME)
	logFile, err := os.Open(logPath)
	utils_.Err(err, "015")
	defer logFile.Close()

	content, err := io.ReadAll(logFile)
	utils_.Err(err, "016")

	fmt.Fprint(os.Stdout, string(content))
}
