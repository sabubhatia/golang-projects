package major

import (
	"rsc.io/quote"
	V3 "rsc.io/quote/v3"
)

func hello() string {
	return quote.Hello()
}

func proverb() string {
	return V3.Concurrency()
}
