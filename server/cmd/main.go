package main

import (
	"fmt"
	"newm/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	if err := app.InitProject(); err != nil {
		panic(fmt.Sprintf(`
		%s%s_________________________________________________________%s
		%s
		_________________________________________________________%s%s
		`, "\n", "\n", "\n", err.Error(), "\n", "\n"))
	}
}
