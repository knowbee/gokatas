package katas_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knowbee/gokatas"
)

var _ = Describe("Katas", func() {
	It("GCAT", func() {
		Expect(DNAtoRNA("GCAT")).To(Equal("GCAU"))
	})

	It("should repeat correctly", func() {
		Expect(RepeatStr(4, "a")).To(Equal("aaaa"))
		Expect(RepeatStr(3, "hello ")).To(Equal("hello hello hello "))
		Expect(RepeatStr(2, "abc")).To(Equal("abcabc"))
	})
	It("should test that the solution returns the correct value", func() {
		Expect(AbbrevName("Sam Harris")).To(Equal("S.H"))
		Expect(AbbrevName("Patrick Feenan")).To(Equal("P.F"))
		Expect(AbbrevName("Evan Cole")).To(Equal("E.C"))
		Expect(AbbrevName("P Favuzzi")).To(Equal("P.F"))
		Expect(AbbrevName("David Mendieta")).To(Equal("D.M"))
	})
	It("should invert number", func() {
		Expect(Opposite(1)).To(Equal(-1))
	})
	It("tests basic strings", func() {
		Expect(IsPalindrome("a")).To(Equal(true))
		Expect(IsPalindrome("aba")).To(Equal(true))
		Expect(IsPalindrome("Abba")).To(Equal(true))
		Expect(IsPalindrome("hello")).To(Equal(false))
	})
	It("should test that the solution returns the correct value", func() {
		Expect(PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
		Expect(PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
		Expect(PositiveSum([]int{})).To(Equal(0))
		Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
		Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
	})
	It("Testing for 1", func() { Expect(NthEven(1)).To(Equal(0)) })
	It("Testing for 2", func() { Expect(NthEven(2)).To(Equal(2)) })
	It("Testing for 3", func() { Expect(NthEven(3)).To(Equal(4)) })
	It("Testing for 100", func() { Expect(NthEven(100)).To(Equal(198)) })
	It("Testing for 1298734", func() { Expect(NthEven(1298734)).To(Equal(2597466)) })
	It("eloquent", func() {
		Expect(RemoveChar("eloquent")).To(Equal("loquen"))
	})
	It("country", func() {
		Expect(RemoveChar("country")).To(Equal("ountr"))
	})
	It("person", func() {
		Expect(RemoveChar("person")).To(Equal("erso"))
	})
	It("place", func() {
		Expect(RemoveChar("place")).To(Equal("lac"))
	})
	It("IsDivisible(3, 3, 4)", func() { Expect(IsDivisible(3, 3, 4)).To(Equal(false)) })
	It("IsDivisible(12, 3, 4)", func() { Expect(IsDivisible(12, 3, 4)).To(Equal(true)) })
	It("IsDivisible(8, 3, 4)", func() { Expect(IsDivisible(8, 3, 4)).To(Equal(false)) })
	It("IsDivisible(48, 3, 4)", func() { Expect(IsDivisible(48, 3, 4)).To(Equal(true)) })
	It("IsDivisible(100, 5, 10)", func() { Expect(IsDivisible(100, 5, 10)).To(Equal(true)) })
	It("IsDivisible(100, 5, 3)", func() { Expect(IsDivisible(100, 5, 3)).To(Equal(false)) })
	It("IsDivisible(4, 4, 2)", func() { Expect(IsDivisible(4, 4, 2)).To(Equal(true)) })
	It("IsDivisible(5, 2, 3)", func() { Expect(IsDivisible(5, 2, 3)).To(Equal(false)) })
	It("IsDivisible(17, 17, 17)", func() { Expect(IsDivisible(17, 17, 17)).To(Equal(true)) })
	It("IsDivisible(17, 1, 17)", func() { Expect(IsDivisible(17, 1, 17)).To(Equal(true)) })
	It("should work for basic tests", func() {
		Expect(CorrectTail("Fox", 'x')).To(BeTrue())
		Expect(CorrectTail("Rhino", 'o')).To(BeTrue())
		Expect(CorrectTail("Meerkat", 't')).To(BeTrue())
	})
	It("should test that the solution returns the correct value", func() {
		Expect(Goals(0, 0, 0)).To(Equal(0))
		Expect(Goals(43, 10, 5)).To(Equal(58))
	})
	It("ExpressionMatter(2, 1, 2)", func() { Expect(ExpressionMatter(2, 1, 2)).To(Equal(6)) })
	It("ExpressionMatter(2, 1, 1)", func() { Expect(ExpressionMatter(2, 1, 1)).To(Equal(4)) })
	It("ExpressionMatter(1, 1, 1)", func() { Expect(ExpressionMatter(1, 1, 1)).To(Equal(3)) })
	It("ExpressionMatter(1, 2, 3)", func() { Expect(ExpressionMatter(1, 2, 3)).To(Equal(9)) })
	It("ExpressionMatter(1, 3, 1)", func() { Expect(ExpressionMatter(1, 3, 1)).To(Equal(5)) })
	It("ExpressionMatter(2, 2, 2)", func() { Expect(ExpressionMatter(2, 2, 2)).To(Equal(8)) })
	It("ExpressionMatter(5, 1, 3)", func() { Expect(ExpressionMatter(5, 1, 3)).To(Equal(20)) })
	It("ExpressionMatter(3, 5, 7)", func() { Expect(ExpressionMatter(3, 5, 7)).To(Equal(105)) })
	It("ExpressionMatter(5, 6, 1)", func() { Expect(ExpressionMatter(5, 6, 1)).To(Equal(35)) })
	It("ExpressionMatter(1, 6, 1)", func() { Expect(ExpressionMatter(1, 6, 1)).To(Equal(8)) })
	It("ExpressionMatter(2, 6, 1)", func() { Expect(ExpressionMatter(2, 6, 1)).To(Equal(14)) })
	It("ExpressionMatter(6, 7, 1)", func() { Expect(ExpressionMatter(6, 7, 1)).To(Equal(48)) })
	It("ExpressionMatter(2, 10, 3)", func() { Expect(ExpressionMatter(2, 10, 3)).To(Equal(60)) })
	It("ExpressionMatter(1, 8, 3)", func() { Expect(ExpressionMatter(1, 8, 3)).To(Equal(27)) })
	It("ExpressionMatter(9, 7, 2)", func() { Expect(ExpressionMatter(9, 7, 2)).To(Equal(126)) })
	It("ExpressionMatter(1, 1, 10)", func() { Expect(ExpressionMatter(1, 1, 10)).To(Equal(20)) })
	It("ExpressionMatter(9, 1, 1)", func() { Expect(ExpressionMatter(9, 1, 1)).To(Equal(18)) })
	It("ExpressionMatter(10, 5, 6)", func() { Expect(ExpressionMatter(10, 5, 6)).To(Equal(300)) })
	It("ExpressionMatter(1, 10, 1)", func() { Expect(ExpressionMatter(1, 10, 1)).To(Equal(12)) })
	It("should test that the solution returns the correct value", func() {
		Expect(MakeNegative(42)).To(Equal(-42))
	})
	It("should return the correct values", func() {
		Expect(MultipleOfIndex([]int{22, -6, 32, 82, 9, 25})).To(ConsistOf(-6, 32, 25))
		Expect(MultipleOfIndex([]int{68, -1, 1, -7, 10, 10})).To(ConsistOf(-1, 10))
		Expect(MultipleOfIndex([]int{11, -11})).To(ConsistOf(-11))
		Expect(MultipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68})).To(ConsistOf(-85, 72, 0, 68))
		Expect(MultipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51})).To(ConsistOf(38, -44, -99))
		Expect(MultipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35})).To(ConsistOf(-49, 8, -60, 35))
	})
	It("should test that the solution returns the correct value", func() {
		Expect(Hero(10, 5)).To(Equal(true))
		Expect(Hero(7, 4)).To(Equal(false))
		Expect(Hero(4, 5)).To(Equal(false))
		Expect(Hero(100, 40)).To(Equal(true))
		Expect(Hero(1500, 751)).To(Equal(false))
		Expect(Hero(0, 1)).To(Equal(false))
	})
	It("26 dogs", func() {
		Expect(HowManyDalmatians(26)).To(Equal("More than a handful!"))
	})
	It("8 dogs", func() {
		Expect(HowManyDalmatians(8)).To(Equal("Hardly any"))
	})
	It("14 dogs", func() {
		Expect(HowManyDalmatians(14)).To(Equal("More than a handful!"))
	})
	It("80 dogs", func() {
		Expect(HowManyDalmatians(80)).To(Equal("Woah that's a lot of dogs!"))
	})
	It("100 dogs", func() {
		Expect(HowManyDalmatians(100)).To(Equal("Woah that's a lot of dogs!"))
	})
	It("50 dogs", func() {
		Expect(HowManyDalmatians(50)).To(Equal("More than a handful!"))
	})
	It("10 dogs", func() {
		Expect(HowManyDalmatians(10)).To(Equal("Hardly any"))
	})
	It("101 dogs", func() {
		Expect(HowManyDalmatians(101)).To(Equal("101 DALMATIONS!!!"))
	})
	examples := [...][2]int{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}
	for i := 0; i < len(examples); i++ {
		v := examples[i]
		It(fmt.Sprintf("Testing for %d", v[0]), func() {
			Expect(Summation(v[0])).To(Equal(v[1]))
		})
	}
	It("Should return the correct positive value", func() {
		Expect(Combat(100.0, 50.0)).To(Equal(50.0))
	})
	It("Should return 0 rather than negative health", func() {
		Expect(Combat(1.0, 100)).To(Equal(0.0))
	})
	It("Does not add a bonus if undeserved", func() {
		Expect(BonusTime(100, false)).To(Equal("£100"))
		Expect(BonusTime(9, false)).To(Equal("£9"))
		Expect(BonusTime(55000, false)).To(Equal("£55000"))
	})
	It("Adds a bonus if deserved", func() {
		Expect(BonusTime(100, true)).To(Equal("£1000"))
		Expect(BonusTime(14000, true)).To(Equal("£140000"))
	})
	It("should test that the solution returns the correct value", func() {
		Expect(Maps([]int{1, 2, 3})).To(Equal([]int{2, 4, 6}))
		Expect(Maps([]int{4, 1, 1, 1, 4})).To(Equal([]int{8, 2, 2, 2, 8}))
		Expect(Maps([]int{2, 2, 2, 2, 2, 2})).To(Equal([]int{4, 4, 4, 4, 4, 4}))
	})
	It("should handle sample Test Cases", func() {
		Expect(OddCount(15)).To(Equal(7))
		Expect(OddCount(15023)).To(Equal(7511))
	})
})
