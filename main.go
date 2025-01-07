package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func startSwayidleService() error {
	command := "systemctl --user start swayidle.service"
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func stopSwayidleService() error {
	command := "systemctl --user stop swayidle.service"
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func getSwayidleServiceStatus() (bool, error) {
	command := "systemctl --user is-active --quiet swayidle.service"
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			/*
				Get exit code to evaluate service status.
				3 == not Running.
				Everything else indicating an Error.
			*/
			if exitError.ExitCode() == 3 {
				return false, nil
			}
		}
		return false, err
	}

	// Der Service läuft
	return true, nil
}

func main() {
	enable := flag.Bool("on", false, "Enable screenlock and display timeout.")
	disable := flag.Bool("off", false, "Disable screenlock and display timeout.")
	flag.Parse()

	// check current status of swayidle service.
	isRunning, err := getSwayidleServiceStatus()
	if err != nil {
		panic(err)
	}

	// start service if not already running.
	if *enable && !isRunning {
		err := startSwayidleService()
		if err != nil {
			fmt.Println("Failed to start swayidle Service!")
			return
		}
	}

	// stop service if not already stoped.
	if *disable && isRunning {
		err := stopSwayidleService()
		if err != nil {
			fmt.Println("Failed to stop swayidle Service!")
			return
		}
	}
}
