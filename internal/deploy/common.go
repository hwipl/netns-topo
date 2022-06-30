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
