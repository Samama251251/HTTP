package parser

import (
	"fmt"
	"strings"
)

const splitter = "\r\n"

type Request struct {
	RequestLine    RequestLine
	RequestHeaders Headers
}
type Headers map[string]string
type RequestLine struct {
	Method        string
	RequestTarget string
	HTTPVersion   string
}

func ParseRequestLine(req string) (*RequestLine, error) {
	fmt.Println("I am in the ParseRequestFunction")
	fmt.Printf("RawRequest:%q", req)
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

func ParseHeaders(req string) (Headers, error) {
	headers := make(Headers)
	fmt.Println("I am in the ParseHeadersFunction")
	lines := strings.Split(req, splitter)
	if len(lines) == 0 || lines[0] == "" {
		return nil, fmt.Errorf("invalid request: empty request line")
	}
	var rawHeaders []string
	for index, header := range lines {
		if index > 0 && index < len(lines)-2 {
			fmt.Println("Appending Header:", header)
			rawHeaders = append(rawHeaders, header)
		}
	}
	for _, rawHead := range rawHeaders {
		parts := strings.SplitN(rawHead, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header line: %q", rawHead)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		headers[strings.ToLower(key)] = value
	}
	fmt.Println("headers:", headers)
	return headers, nil

}
