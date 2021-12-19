package demonizer

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"
)

const BLOCK_ENV_VAR = "BLOCK_ENV_VAR"
const DEMONIZED_PID = "DEMONIZED_PID"

//Creates new d-process using
//passed command
func DemonizeProcess(command string, args ...string) {
	if _, ok := os.LookupEnv(BLOCK_ENV_VAR); !ok {
		cmd := exec.Command(command, args...)
		cmd.Env = append(cmd.Env, os.Args...)
		cmd.Env = append(cmd.Env, BLOCK_ENV_VAR)
		if err := cmd.Start(); err != nil {
			logrus.Fatal()
		}
		os.Exit(0)
	}
	syscall.Umask(022)
	workPath, err := filepath.Abs(command)
	if err != nil {
		logrus.Fatal()
	}

	if err := os.Chdir(workPath); err != nil {
		logrus.Fatal()
	}

	if _, err := syscall.Setsid(); err != nil {
		logrus.Fatal()
	}
}

//Demonizes this executable process
func DemonizeThisProcess() {
	DemonizeProcess(os.Args[0])
}
