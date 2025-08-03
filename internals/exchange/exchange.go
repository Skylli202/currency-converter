// Package exchange
package exchange

type Exchange interface {
	Convert(amt float64, from, to string) float64
}
