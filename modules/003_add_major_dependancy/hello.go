package major

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func hello() string {
	return quote.Hello()
}

func proverb() string {
	return quoteV3.Concurrency()
}
