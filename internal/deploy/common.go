package deploy

import (
	"log"
	"os"
	"os/exec"
)

// runIP runs the ip command with the parameters params
func runIP(params ...string) {
	cmd := exec.Command("ip", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
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
