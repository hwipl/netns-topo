package deploy

import (
	"io"
	"log"
	"os"
	"os/exec"
)

// runIPStdinOutErr runs the ip command with Stdin, Stdout, Stderr and parameters params
func runIPStdinOutErr(stdin io.Reader, stdout, stderr io.Writer, params ...string) {
	cmd := exec.Command("ip", params...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

// runIP runs the ip command with the parameters params
func runIP(params ...string) {
	runIPStdinOutErr(nil, os.Stdout, os.Stderr, params...)
}

// runNetns runs the command and its parameters in params in netns
func runNetns(netns string, params ...string) {
	p := []string{"netns", "exec", netns}
	p = append(p, params...)
	runIP(p...)
}

// runNetnsIP runs the ip command in netns with the parameters params
func runNetnsIP(netns string, params ...string) {
	p := []string{"ip"}
	p = append(p, params...)
	runNetns(netns, p...)
}
