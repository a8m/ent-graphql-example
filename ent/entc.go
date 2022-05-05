//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// The codegen is executed from generate.go.
	// So the path for the config file, ent schema, and the GQL schema
	// starts from the project root.

	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		// Generate GQL schema from the Ent's schema.
		entgql.WithSchemaGenerator(),
		// Generate the filters to a separate schema
		// file and load it in the gqlgen.yml config.
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereFilters(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./ent/template"),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
