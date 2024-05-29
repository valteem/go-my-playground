package examine

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/coreos/go-systemd/activation"
)

func myExampleCmd(binaryName string) (string, []string) {
	sourcePath := fmt.Sprintf("../examples/activation/%s.go", binaryName)
	sourceCmdLine := []string{"go", "run", sourcePath}
	binaryPath := fmt.Sprintf("../test_bins/%s.example", binaryName)
	if _, err := os.Stat(binaryPath); err != nil && os.IsNotExist(err) {
		return sourceCmdLine[0], sourceCmdLine[1:]
	}
	return binaryPath, []string{binaryPath}
}

func TestListeners(t *testing.T) {

	arg0, cmdline := myExampleCmd("listen")
	cmd := exec.Command(arg0, cmdline...)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "LISTEN_FDS=2", "LISTEN_FDNAMES=fd1:fd2", "FIX_LISTEN_PID=1")

	_, e := activation.Listeners()
	if e != nil {
		t.Errorf("Listeners(): get %e error", e)
	}
	// if len(l) == 0 {
	// 	t.Errorf("Listeners() returned nothing")
	// }

}
