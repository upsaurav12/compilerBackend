package executor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Execute(language string, code []byte, input string) (string, string) {

	tempDir, err := ioutil.TempDir("", "compiler")
	if err != nil {
		return "", "Failed to create Temp dir"
	}
	defer os.RemoveAll(tempDir)

	var fileName string
	var dockerFiles string
	switch language {
	case "go":
		fileName = "main.go"
		dockerFiles = "dockerfiles/go/Dockerfile"
		err = os.WriteFile(tempDir+"/"+fileName, code, 0644)
		if err != nil {
			return "", "Error writing Go code to file"
		}
	case "cpp":
		fileName = "main.cpp"
		dockerFiles = "dockerfiles/cpp/Dockerfile"
		err = os.WriteFile(tempDir+"/"+fileName, code, 0644)
		if err != nil {
			return "", "Error while writing cpp code to the file"
		}
	default:
		return "", "Unsupported language"
	}

	output, execErr := runDockerContainer(tempDir, dockerFiles, fileName, input)
	if execErr != nil {
		return "", fmt.Sprintf("Error executing code: %v", execErr)
	}

	return output, ""
}

func runDockerContainer(tempDir string, dockerFiles string, fileName string, input string) (string, error) {
	cmd := exec.Command("sudo", "docker", "build", "-t", "compiler-"+fileName, "-f", dockerFiles, tempDir)
	var buildOut, buildErr bytes.Buffer
	cmd.Stdout = &buildOut
	cmd.Stderr = &buildErr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Docker build failed: %s", buildErr.String())
	}

	// Run the Docker container
	cmd = exec.Command("sudo", "docker", "run", "--rm", "compiler-"+fileName)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}

	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Error running Docker container: %s", errOut.String())
	}

	return out.String(), nil
}
