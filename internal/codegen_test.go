package internal

import (
	"fmt"
	"os"
	"testing"

	"nimona.io"
)

func TestCodegen(t *testing.T) {
	if os.Getenv("NIMONA_CODEGEN") == "" {
		t.Skip("skipping codegen")
	}

	fmt.Println("generating nimona document methods")

	err := nimona.GenerateDocumentMethods(
		"models_gen.go",
		"internal",
		NimonaFeed{},
		NimonaProfile{},
		NimonaNote{},
		NimonaFollow{},
	)
	if err != nil {
		panic(fmt.Errorf("error generating nimona document methods: %w", err))
	}

	fmt.Println("done")
}
