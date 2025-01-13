// Copyright 2024 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The goal of this wrapper is to detect panic that indicates corruption of checkpoints folder
// (which can happen during kernel crash or system restart) and in case it is detected
// it removes the checkpoint folder.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const (
	panicMessage = "panic: assertion failed: Page expected to be"
)

func main() {
	checkpointDir := os.Getenv("CHECKPOINT_DIR")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Missing command arguments")
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout // Redirect stdout of cmd to stdout of wrapper

	r, w := io.Pipe()
	cmd.Stderr = w // Redirect stderr of cmd to the writer end of the pipe

	var wg sync.WaitGroup
	wg.Add(1)

	// Goroutine to monitor the stderr output for panic message
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(r) // Read from the reader end of the pipe
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintln(os.Stderr, line) // Forward the stderr output to stderr of wrapper
			if strings.Contains(line, panicMessage) {
				fmt.Fprintln(os.Stderr, "Specific panic detected, deleting all files in checkpoint folder...")
				if checkpointDir != "" {
					err := filepath.Walk(checkpointDir, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if path != checkpointDir { // Skip the root directory
							return os.RemoveAll(path)
						}
						return nil
					})
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error:", err)
					}
				}
			}
		}
	}()

	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Wait for the command to exit
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	// Close the writer end of the pipe to signal the end of data
	w.Close()

	// Wait for the output monitoring goroutine to finish
	wg.Wait()
}
