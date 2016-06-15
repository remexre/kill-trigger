package agora

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/agora/runtime"
)

// PastebinResolver resolves modules from Pastebin.
type PastebinResolver struct {
	client *http.Client
}

// NewPastebinResolver creates a new PastebinResolver.
func NewPastebinResolver() *PastebinResolver {
	return &PastebinResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Resolve resolves a module.
func (r *PastebinResolver) Resolve(mod string) (io.Reader, error) {
	res, err := r.client.Get("http://pastebin.com/raw/" + mod)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, runtime.NewModuleNotFoundError(mod)
	} else if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %s", res.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	return buf, err
}
