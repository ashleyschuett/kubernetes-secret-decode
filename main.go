package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	// dataKey is the name of the key that is populated with the Kubernetes
	// secret's data
	dataKey = "data"
	// stringDataKey is the name of the key that is used writing a secret's plain
	// text information to save to the secret. On save Kubernetes encodes this field
	stringDataKey = "stringData"
)

var outputType string

// SecretData extracts out the data portion of a Kubernetes secret
type SecretData struct {
	Data map[string]string `json:"data" yaml:"data"`
}

// Secret allows us to read and return the full Kubernetes secret
type Secret map[string]interface{}

// Unmarshallable allows me to unmarsal different strings with the same interface
type Unmarshallable func([]byte, interface{}) error

func main() {
	// parse the rest of the command to kubectl and run it
	output := parseAndRunCommand()
	unmarshal := getUnmarshalByOutputType(outputType)

	sd, err := getDecodedSecretData(unmarshal, output.Bytes())
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	s, err := getFullSecretWithDecodedData(unmarshal, output.Bytes(), sd)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	secret := getStringSecret(s, outputType)
	// Print exposed secret
	fmt.Fprintf(os.Stdout, "%s\n", secret)
}

func parseAndRunCommand() bytes.Buffer {
	args := getKubectlArgs()

	// check that the global output type was set, if it's not set we can not decode the secret
	if outputType == "" {
		fmt.Fprintf(os.Stdout, "please set -o flag to json or yaml\n")
		os.Exit(1)
	}

	cmd := exec.Command("kubectl", args...)

	var output, errb bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stdout, errb.String())
		os.Exit(1)
	}

	return output
}

func getKubectlArgs() []string {
	var args []string

	for i, arg := range os.Args {
		// remove the ksd binary from the kubectl command
		if strings.Contains(arg, "kubectl-ksd") || strings.Contains(arg, "kubernetes-secret-decode") {
			args = append(os.Args[:i], os.Args[i+1:]...)
		}

		if strings.Contains(arg, "json") || strings.Contains(arg, "yaml") {
			// this set the global variable that is used for parsing the output
			outputType = strings.Trim(arg, "-o")
		}
	}

	return args
}

func getUnmarshalByOutputType(outputType string) Unmarshallable {
	var unmarshal Unmarshallable
	if isJSON(outputType) {
		unmarshal = json.Unmarshal
	} else {
		unmarshal = yaml.Unmarshal
	}

	return unmarshal
}

func getStringSecret(s *Secret, outputType string) string {
	var secret []byte
	if isJSON(outputType) {
		secret, _ = json.MarshalIndent(s, "", "    ")
	} else {
		secret, _ = yaml.Marshal(s)
	}

	return string(secret)
}

func getFullSecretWithDecodedData(unmarshal Unmarshallable, output []byte, sd *SecretData) (*Secret, error) {
	var s Secret
	var err error

	err = unmarshal(output, &s)
	if err != nil {
		return nil, err
	}

	for key := range s {
		if key == dataKey {
			s[stringDataKey] = sd.Data
		}
	}

	delete(s, dataKey)

	return &s, nil
}

func getDecodedSecretData(unmarshal Unmarshallable, output []byte) (*SecretData, error) {
	var s SecretData
	var err error

	err = unmarshal(output, &s)
	if err != nil {
		return nil, err
	}

	err = parseData(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func getKubectlSecretOutput() []byte {
	var output []byte
	reader := bufio.NewReader(os.Stdout)

	for {
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		}

		output = append(output, input)
	}

	return output
}

func isJSON(o string) bool {
	return o == "json"
}

func parseData(s *SecretData) error {
	var err error
	for key, value := range s.Data {
		s.Data[key], err = decodeString(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeString(encoded string) (string, error) {
	decoded, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
