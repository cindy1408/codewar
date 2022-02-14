package main

import (
	// . "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(a1 []string, a2 []string, exp int) {
	var ans = MxDifLg(a1, a2)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests MxDifLg", func() {

	It("should handle basic cases", func() {
		var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
		var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
		dotest(s1, s2, 13)
		s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
		s2 = []string{"bbbaaayddqbbrrrv"}
		dotest(s1, s2, 10)
	})
})
