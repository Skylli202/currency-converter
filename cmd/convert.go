package cmd

import (
	"errors"
	"fmt"
	"strconv"

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
	Args: ValidateConvertCmdArgs,
	RunE: ConvertCmdRunE,
}

func ValidateConvertCmdArgs(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 0:
		return nil
	case 1:
		_, err := strconv.ParseFloat(args[0], 64)
		return err
	case 2:
		// format <?> <?>
		_, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			_, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				// format <NaN> <NaN> --> return err because none of the args is a float
				return err
			}
			// format <TARGET_CURR> <AMOUNT>
			if !IsValidCurrencyCode(args[0]) {
				return fmt.Errorf("error: %s is not a valid currency code", args[0])
			}
			return nil
		}
		// format <AMOUNT> <TARGET_CURR>
		if !IsValidCurrencyCode(args[1]) {
			return fmt.Errorf("error: %s is not a valid currency code", args[1])
		}
		return nil
	case 3:
		// format <?> <?> <?>
		_, err0 := strconv.ParseFloat(args[0], 64)
		_, err2 := strconv.ParseFloat(args[2], 64)
		if err0 != nil && err2 != nil {
			// format <Nan> <?> <NaN> --> return err because neither the first nor last arguments are float.
			return fmt.Errorf("error: %s is not a valid currency code", args[0])
		} else if err0 != nil {
			// format <FROM_CURR> <TARGET_CURR> <AMOUNT>
			if !IsValidCurrencyCode(args[0]) {
				return fmt.Errorf("error: %s is not a valid currency code", args[1])
			}
			if !IsValidCurrencyCode(args[1]) {
				return fmt.Errorf("error: %s is not a valid currency code", args[1])
			}
			return nil
		} else if err2 != nil {
			// format <AMOUNT> <FROM_CURR> <TARGET_CURR>
			if !IsValidCurrencyCode(args[1]) {
				return fmt.Errorf("error: %s is not a valid currency code", args[1])
			}
			if !IsValidCurrencyCode(args[2]) {
				return fmt.Errorf("error: %s is not a valid currency code", args[1])
			}
			return nil
		} else {
			// format <AMOUNT> <?> <AMOUNT>
			return fmt.Errorf("error: you cannot provide two amonts")
		}
	default:
		return errors.New("too many arguments")
	}
}

// IsValidCurrencyCode return true if the given string `code` is a currency
// code like "USD", "CAD", etc.
func IsValidCurrencyCode(code string) bool {
	// HACK: clearly not a real validation but that's what it is.
	return len(code) == 3
}

func ConvertCmdRunE(cmd *cobra.Command, args []string) error {
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
