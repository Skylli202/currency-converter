package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert [amount] [from] [to]",
	Short: "convert currencies",
	Long: `Convert an amount from a currency to another currency.
The commmand can be called without any argument. The command will prompt the user for
the amount and the from/to currencies. The command also accept argument(s). If one
argument is provided then it is expected to be the amount. If two arguments are 
provided then it is expected to have one number (the amount) and one currency code
(the to currency). The from currency will default to "USD". if three arguments are
provided then it is expected to either have, in that order, "amt from to" or "from to amt".`,
	Example: `// You will be prompted to input the amount and the currenecies.
currency-converter convert

// You will be prompted to input the currencies.
currency-converter convert 1234

// You will NOT be prompted for the from currency, it will use "USD" as default.
currency-converter convert 1234 GBP

// You will NOT be prompted at all.
// When specifying all three arguments the order can either be "amount from to" or "from to amount".
currency-converter convert 1234 USD CAD
currency-converter convert EUR GBP 1234`,
	RunE: ConvertRunE,
}

func ConvertRunE(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented")
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
