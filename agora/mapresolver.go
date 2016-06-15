package agora

import (
	"errors"
	"io"
	"strings"
)

// MapResolver resolves modules based on the contents of a map of strings.
type MapResolver map[string]string

// Resolve resolves a module.
func (m MapResolver) Resolve(mod string) (io.Reader, error) {
	src, ok := m[mod]
	if !ok {
		return nil, errors.New("cannot find module")
	}
	return strings.NewReader(src), nil
}
