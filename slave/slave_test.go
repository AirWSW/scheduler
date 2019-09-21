package slave

import (
	// "errors"
	"testing"
	"strings"
	"fmt"
	"os/exec"
	// "net/http"
	// "time"

	"github.com/AirWSW/scheduler/config"
)

func Test_runner(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		// return nil, err
	}

	args := strings.Split(cfg.Schedule.Tasks[0].Command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	res, err := cmd.Output()
	if err != nil {
		args = strings.Split(cfg.Schedule.Tasks[0].ExecOnFailure, " ")
		cmd = exec.Command(args[0], args[1:]...)
		res, err = cmd.Output()
		fmt.Println(string(res))
	}else{
		fmt.Println(string(res))
		args = strings.Split(cfg.Schedule.Tasks[0].ExecOnSuccess, " ")
		cmd = exec.Command(args[0], args[1:]...)
		res, err = cmd.Output()
		fmt.Println(string(res))
	}
}
