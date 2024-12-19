package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func killSwayidleProcess() error {
	command := "ps -ef | grep swayidle | grep -v grep | grep -v ctrl"
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	var pids []string
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 {
			pids = append(pids, fields[1])
		}
	}

	// Beenden der Prozesse
	for _, pidStr := range pids {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			return err
		}

		proc, err := os.FindProcess(pid)
		if err != nil {
			continue
		}

		err = proc.Kill()
		if err != nil {
			return err
		}
	}
	return nil
}

func createSwayidleProcess(t1, t2 int) error {
	command := fmt.Sprintf("swayidle -w timeout %d 'swaylock -f -c 3B4252' timeout %d 'swaymsg \"output * power off\"' resume 'swaymsg \"output * power on\"' before-sleep 'swaylock -f -c 3B4252'", t1, t2)
	cmd := exec.Command("sh", "-c", command)

	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	t1 := flag.Int("t1", 300, "Set timeout for screenlock.")
	t2 := flag.Int("t2", 600, "Set timeout for display.")
	disable := flag.Bool("off", false, "Disable screenlock and display timeout.")

	flag.Parse()

	// kill running swayidle instance.
	err := killSwayidleProcess()
	if err != nil {
		panic(err)
	}

	if !*disable {
		err := createSwayidleProcess(*t1, *t2)
		if err != nil {
			panic(err)
		}
	}
}
