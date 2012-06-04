/*
Package jsonutil implements utility functions for working with json data.
*/
package jsonutil

import (
	"encoding/json"
	"io"
	"os"
)

// writes out the struct v to the provided writer w.
func Write(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

// indents the struct v and prints it to the writer w.  This function causes
// the entire structure to be encoded into memory before being written
// into w.
func WriteIndented(w io.Writer, v interface{}, prefix, indent string) error {
	raw, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return err
	}
    _, err = w.Write(raw)
    return err
}

// prints the json structure to stdout.
func Print(v interface{}) error {
	return Write(os.Stdout, v)
}

// pretty prints the json structure to stdout.  This function causes the entire
// struct to be encoded into memory before being written to stdout.
func PrettyPrint(v interface{}) error {
	return WriteIndented(os.Stdout, v, "", "  ")
}

func Split(r io.Reader, c chan *json.RawMessage, e chan error) {
    defer close(c)
    defer close(e)
    d := json.NewDecoder(r)
    for {
        var raw json.RawMessage
        err := d.Decode(&raw)
        switch err {
        case nil:
            c <- &raw
        case io.EOF:
            return
        default:
            e <- err
            return
        }
    }
}
