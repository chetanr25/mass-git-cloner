package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptUsername prompts the user to enter a GitHub username or organization
func PromptUsername() (string, error) {
	fmt.Print("Enter GitHub username or organization: ")

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

// PromptConfirmation prompts the user for a yes/no confirmation
func PromptConfirmation(message string) (bool, error) {
	fmt.Printf("%s (y/N): ", message)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	response = strings.TrimSpace(strings.ToLower(response))
	// Default to "N" - only return true for explicit "y" or "yes"
	return response == "y" || response == "yes", nil
}

// PromptRetry asks if the user wants to retry an operation
func PromptRetry(operation string) (bool, error) {
	return PromptConfirmation(fmt.Sprintf("Failed to %s. Would you like to retry?", operation))
}

// DisplayWelcome displays a welcome message
func DisplayWelcome() {
	fmt.Println("🚀 Mass Git Cloner")
	fmt.Println("==================")
	fmt.Println("Clone multiple repositories from a GitHub user or organization")
	fmt.Println()
}

// DisplayError displays an error message
func DisplayError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
}

// DisplaySuccess displays a success message
func DisplaySuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

// DisplayInfo displays an info message
func DisplayInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}
