package uml

import (
	"fmt"
	"io"
	"os/exec"
)

func UML(src []byte, args ...string) (dist []byte, err error) {
	var path string
	var stdin io.WriteCloser

	path, err = exec.LookPath("plantuml")
	if err != nil {
		return
	}

	params := append([]string{"-Tsvg", "-p", "-Djava.awt.headless=true"}, args...)
	cmd := exec.Command(path, params...)
	cmd.Env = append(cmd.Env, "JAVA_OPTS='-Djava.awt.headless=true'")

	stdin, err = cmd.StdinPipe()
	if err != nil {
		return
	}

	_, err = stdin.Write(src)
	if err != nil {
		return
	}

	stdin.Close()
	if err != nil {
		return
	}

	dist, err = cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf(string(dist))
		dist = []byte{}
	}

	return
}
