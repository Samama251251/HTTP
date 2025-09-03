package parser

import (
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *RequestLine
		expectErr bool
	}{
		{
			name:  "valid GET request",
			input: "GET /hello HTTP/1.1\r\n",
			want: &RequestLine{
				Method:        "GET",
				RequestTarget: "/hello",
				HTTPVersion:   "1.1",
			},
			expectErr: false,
		},
		{
			name:      "empty request line",
			input:     "\r\n",
			want:      nil,
			expectErr: true,
		},
		{
			name:      "missing version",
			input:     "GET /hello\r\n",
			want:      nil,
			expectErr: true,
		},
		{
			name:      "invalid HTTP version format",
			input:     "GET /hello HTTP\r\n",
			want:      nil,
			expectErr: true,
		},
		{
			name:  "extra spaces between fields",
			input: "GET    /hello    HTTP/1.0\r\n",
			want: &RequestLine{
				Method:        "GET",
				RequestTarget: "/hello",
				HTTPVersion:   "1.0",
			},
			expectErr: false,
		},
		{
			name:  "different method (POST)",
			input: "POST /submit HTTP/2.0\r\n",
			want: &RequestLine{
				Method:        "POST",
				RequestTarget: "/submit",
				HTTPVersion:   "2.0",
			},
			expectErr: false,
		},
		{
			name:      "no request line at all",
			input:     "",
			want:      nil,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseRequestLine(tc.input)

			if tc.expectErr && err == nil {
				t.Fatalf("expected error but got none (input=%q)", tc.input)
			}
			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectErr {
				if got.Method != tc.want.Method ||
					got.RequestTarget != tc.want.RequestTarget ||
					got.HTTPVersion != tc.want.HTTPVersion {
					t.Errorf("got %+v, want %+v", got, tc.want)
				}
			}
		})
	}
}
