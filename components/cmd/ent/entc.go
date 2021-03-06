// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
		Target:    "./generated",
		Package:   "github.com/autoai-org/aid/ent/generated",
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
