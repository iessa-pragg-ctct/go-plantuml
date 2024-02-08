package uml

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func UML(src []byte, args ...string) (dist []byte, err error) {
	var path string
	var stdin io.WriteCloser

	path, err = exec.LookPath("plantuml")
	if err != nil {
		log.Println("could not find plantuml in path")
		log.Println(err)
		return
	}

	params := append([]string{"-Tsvg", "-p"}, args...)
	cmd := exec.Command(path, params...)
	cmd.Env = append(os.Environ(), "JAVA_OPTS='-Djava.awt.headless=true'")

	stdin, err = cmd.StdinPipe()
	if err != nil {
		log.Println("failed to set up stdin pipe")
		log.Println(err)
		return
	}

	_, err = stdin.Write(src)
	if err != nil {
		log.Println("failed to write to stdin")
		log.Println(err)
		return
	}

	stdin.Close()
	if err != nil {
		log.Println("failed to close stdin")
		log.Println(err)
		return
	}

	dist, err = cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to run plantuml")
		log.Println(err)
		err = fmt.Errorf(string(dist))
		dist = []byte{}
	}

	return
}
