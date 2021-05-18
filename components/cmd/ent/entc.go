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
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
