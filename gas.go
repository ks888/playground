package main

import "os/exec"

func main() {
	_ = exec.Command("ls") // #nosec
	_ = []*exec.Cmd{
		exec.Command("ls"),
	} // #nosec
	_ = struct{ cmd *exec.Cmd }{
		exec.Command("ls"), // #nosec
	}
}
