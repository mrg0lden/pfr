package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/theckman/yacspin"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Torrent URL wasn't provided")
		os.Exit(1)
	}
	url := os.Args[1]
	spinner := prepareSpinner()

	pf := exec.Command("go-peerflix", url)
	mpv := exec.Command("mpv", "http://:8080")

	pf.Start()
	spinner.Message("Launching PeerFlix")
	spinner.Start()
	spinner.StopMessage("Ready to play")
	time.Sleep(time.Second * 3)
	spinner.Stop()
	mpv.Start()

	fmt.Println("Enjoy ✨")
	spinner.StopMessage("All cleaning done.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for range sig {
			cleanup(spinner, pf)
		}
	}()
	mpv.Wait()
	cleanup(spinner, pf)
}

func cleanup(spinner *yacspin.Spinner, pf *exec.Cmd) {
	fmt.Println("I hope you enjoyed.")

	spinner.Start()

	spinner.Message("Shutting down peerflix")

	pf.Process.Kill()
	pf.Wait()

	spinner.Stop()
	os.Exit(1)
}

func prepareSpinner() *yacspin.Spinner {
	cfg := yacspin.Config{
		Frequency:     time.Second / 10,
		CharSet:       yacspin.CharSets[26],
		StopCharacter: "✔",
		StopColors:    []string{"fgGreen"},
		Message:       "Cleaning up started",
		StopMessage:   "All done.",
	}
	spinner, err := yacspin.New(cfg)
	if err != nil {
		panic(err)
	}
	return spinner
}
