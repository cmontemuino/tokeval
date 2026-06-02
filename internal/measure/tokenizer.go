package measure

import "context"

// Tokenizer defines the interface for counting tokens in a byte stream.
type Tokenizer interface {
	// Name returns the identifier of the tokenizer (e.g., "openai:cl100k_base").
	Name() string

	// count returns the number of tokens in the given context.
	Count(ctx context.Context, content []byte) (int, error)
}
