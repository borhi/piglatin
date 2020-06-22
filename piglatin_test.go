package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var encodeTests = []struct {
	in  string
	out string
}{
	{
		in:  "pig latin banana will butler happy duck me honest",
		out: "igpay atinlay ananabay illway utlerbay appyhay uckday emay onesthay",
	},
	{
		in:  "smile string stupid glove trash floor store",
		out: "ilesmay ingstray upidstay oveglay ashtray oorflay orestay",
	},
	{
		in:  "eat omelet are egg explain always ends",
		out: "eatyay omeletyay areyay eggyay explainyay alwaysyay endsyay",
	},
}

func TestEncode(t *testing.T) {
	for i, test := range encodeTests {
		actual := Encode(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}
}
