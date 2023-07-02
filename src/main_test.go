package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set up any necessary test environment
	setup()

	// Run the tests
	code := m.Run()

	// Teardown any resources if required
	teardown()

	// Exit with the test result code
	os.Exit(code)
}

func setup() {
	// Perform setup tasks if any
	// For example, set up environment variables required for testing
	os.Setenv("SERVER_VERSION", "1.1")
	os.Setenv("SERVER_SERVICE_NAME", "example-server")
	os.Setenv("SERVER_HTTP_PORT", "2001")
	os.Setenv("LOG_LEVEL", "debug")
}

func teardown() {
	// Perform teardown tasks if any
}

func TestMain_StartsServer(t *testing.T) {
	// Mock any dependencies or perform any necessary setup

	// Capture the log output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Call the main function
	go main()

	// Wait for the server to start (you may need to introduce a delay here if needed)

	// Make a request to the server
	resp, err := http.Get("http://localhost:2001")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Stop the server gracefully (if applicable)
	// Send a signal to terminate the server gracefully
	//err = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	//assert.NoError(t, err)

	// Check the log output
	logOutput := buf.String()
	assert.Contains(t, logOutput, "server started...")
}

func TestMain_UsesDefaultValues(t *testing.T) {
	// Mock any dependencies or perform any necessary setup

	// Unset environment variables to test default values
	os.Unsetenv("SERVER_VERSION")
	os.Unsetenv("SERVER_SERVICE_NAME")
	os.Unsetenv("SERVER_HTTP_PORT")
	os.Unsetenv("LOG_LEVEL")

	// Capture the log output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Call the main function
	go main()

	// Wait for the server to start (you may need to introduce a delay here if needed)

	// Make a request to the server using default port
	resp, err := http.Get("http://localhost:2000")
	assert.NoError(t, err)
	log.Info(err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Stop the server gracefully (if applicable)
	// Send a signal to terminate the server gracefully

	// Check the log output
	logOutput := buf.String()
	assert.Contains(t, logOutput, "server started...")
}

// Add more test cases as needed
