// Package front provides YAML frontmatter unmarshalling.
package frontmattergo

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// Delimiter.
var delimiter = []byte("---\n")

// Unmarshal YAML frontmatter.
func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	// no frontmatter
	if !bytes.HasPrefix(b, delimiter) {
		return b, nil
	}

	// split into 3 parts
	parts := bytes.SplitN(b, delimiter, 3)

	content = parts[2]

	if err = yaml.Unmarshal(parts[1], v); err != nil {
		return
	}

	return
}
