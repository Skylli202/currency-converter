/*
Copyright © 2025 Elouan GOUINGUENET <elouangouinguenet@gmail.com>
*/
package main

import (
	"fmt"
	"os"

	"github.com/Skylli202/currency-converter/internals/exchange"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error: could not load env; %s", err.Error())
	}
	// cmd.Execute()
	appID := os.Getenv("APP_ID")
	ex := exchange.NewOpenexchangeratesExchange(appID)
	amt := 24.
	cAmt := ex.Convert(amt, "USD", "CAD")
	fmt.Printf("USD: %.2f\nCAD: %.2f\n", amt, cAmt)
}
