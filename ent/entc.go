// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
	}, entc.TemplateDir("./template"))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
