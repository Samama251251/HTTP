package parser

import (
	"fmt"
	"strings"
)

const splitter = "\r\n"

type Request struct {
	RequestLine RequestLine
}
type RequestLine struct {
	Method        string
	RequestTarget string
	HTTPVersion   string
}

func ParseRequestLine(req string) (*RequestLine, error) {
	fmt.Println("I am in the ParseRequestFunction")

	lines := strings.Split(req, splitter)
	if len(lines) == 0 || lines[0] == "" {
		return nil, fmt.Errorf("invalid request: empty request line")
	}

	// Here we will split the request line into its individual parts
	parts := strings.Fields(lines[0])
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %q", lines[0])
	}

	method := parts[0]
	target := parts[1]
	rawVersion := parts[2]
	versionParts := strings.Split(rawVersion, "/")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("invalid HTTP version: %q", rawVersion)
	}
	version := versionParts[1]

	requestLine := &RequestLine{
		Method:        method,
		RequestTarget: target,
		HTTPVersion:   version,
	}

	return requestLine, nil
}
