package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// LoadJSON loads and unmarshals a specified JSON file into a supplied structure pointer.
func LoadJSON(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}

// SaveJSON marshals, indents and saves a supplied structure to the specified file.
func SaveJSON(filename string, src interface{}) error {
	var data []byte
	var err error
	data, err = json.MarshalIndent(src, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0600)
}

// JSONCmd prints data structures for advanced input.
type JSONCmd struct {
	CP JSONCPCmd `command:"counterparty" alias:"cp" description:"Print the input JSON for external counterparties."`
}

// JSONCPCmd prints an empty ExternalCounterparty structure.
type JSONCPCmd struct{}

func (cmd *JSONCPCmd) Execute(args []string) error {
	var cp revolut.ExternalCounterparty
	s, err := json.MarshalIndent(&cp, "", "\t")
	if err != nil {
		return err
	}
	slog.Msg("%s", s)
	return nil
}
