package agora

import (
	"io"
	"regexp"

	"github.com/PuerkitoBio/agora/runtime"
)

type aggResEntry struct {
	Prefix   string
	Resolver runtime.ModuleResolver
}

var aggResRegex = regexp.MustCompile(`(.*?):(.*)`)

// An AggregateResolver holds several resolvers, and looks them up by prefix, or
// chooses a default if the module is unprefixed.
type AggregateResolver struct {
	defaultResolver runtime.ModuleResolver
	entries         []aggResEntry
}

// NewAggregateResolver creates a new AggregateResolver.
func NewAggregateResolver(defaultResolver runtime.ModuleResolver) *AggregateResolver {
	return &AggregateResolver{defaultResolver: defaultResolver}
}

// Add adds a new Resolver.
func (r *AggregateResolver) Add(prefix string, resolver runtime.ModuleResolver) {
	r.entries = append(r.entries, aggResEntry{prefix, resolver})
}

// Resolve resolves a module.
func (r *AggregateResolver) Resolve(mod string) (io.Reader, error) {
	matches := aggResRegex.FindStringSubmatch(mod)
	if len(matches) == 3 && matches[0] == mod {
		for _, entry := range r.entries {
			if entry.Prefix == matches[1] {
				return entry.Resolver.Resolve(matches[2])
			}
		}
	}

	return r.defaultResolver.Resolve(mod)
}
