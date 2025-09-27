package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptUsername() (string, error) {
	fmt.Print("\033[33mEnter GitHub username or organization: \033[0m")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	username = strings.TrimSpace(username)
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return username, nil
}

func PromptConfirmation(message string) bool {
	fmt.Printf("%s (y/N): ", message)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

func PromptRetry(operation string) bool {
	return PromptConfirmation(fmt.Sprintf("Failed to %s. Would you like to retry?", operation))
}

func DisplayWelcome() {
	fmt.Println(`                           
  ____ ___ _____     ____ _     ___  _   _ _____ ____  
 / ___|_ _|_   _|   / ___| |   / _ \| \ | | ____|  _ \ 
| |  _ | |  | |    | |   | |  | | | |  \| |  _| | |_) |
| |_| || |  | |    | |___| |__| |_| | |\  | |___|  _ < 
 \____|___| |_|     \____|_____\___/|_| \_|_____|_| \_\`)

	// fmt.Println("\n\n==================")
	fmt.Println("\n\n\033[32mClone multiple repositories from a GitHub user or organization\033[0m")
	fmt.Println()
}

func DisplayError(err error) {
	fmt.Printf("Error: %v\n", err)
}

func DisplaySuccess(message string) {
	fmt.Printf("%s\n", message)
}

func DisplayInfo(message string) {
	fmt.Printf("\033[36m%s\033[0m\n", message)
}
