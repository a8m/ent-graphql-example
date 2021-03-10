// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/contrib/entgql"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}