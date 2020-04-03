package katas

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// DNAtoRNA https://www.codewars.com/kata/5556282156230d0e5e000089/train/go
// DNAtoRNA DNA to RNA Conversion
func DNAtoRNA(dna string) string {

	var newsequence string = ""
	for _, letter := range dna {
		if string(letter) == "T" {
			newsequence += "U"
		} else {

			newsequence += string(letter)
		}
	}
	return newsequence
}

// GetSize https://www.codewars.com/kata/565f5825379664a26b00007c/train/go
// GetSize Surface Area and Volume of a Box
func GetSize(w, h, d int) [2]int {
	// your code here
	volume := 1 * w * h * d
	area := 2*h*w + 2*h*d + 2*w*d
	res := [2]int{area, volume}
	// res[0] = area
	// res[1] = volume
	return res
}

// BoolToWord https://www.codewars.com/kata/53369039d7ab3ac506000467/train/go
// BoolToWord Convert boolean values to strings 'Yes' or 'No'.
func BoolToWord(word bool) string {

	if word {
		return "Yes"
	}
	return "No"
}

// RepeatStr https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
// RepeatStr repeatStr(6, "I") // "IIIIII"
// RepeatStr repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"

// RepeatStr String repeat
func RepeatStr(repetitions int, value string) string {
	// res := ""
	// for i := 1; i <= repetitions; i++ {
	// 	res += value
	// }
	// return res
	return strings.Repeat(value, repetitions)
}

// AbbrevName https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go
// AbbrevName Abbreviate a Two Word Name
func AbbrevName(name string) string {
	//your code here
	sep := " "
	pieces := strings.Split(name, sep)
	return strings.ToUpper(string(string(pieces[0][0]) + string(".") + string(pieces[1][0])))
}

// Opposite https://www.codewars.com/kata/56dec885c54a926dcd001095/train/go
// Opposite number
func Opposite(value int) int {
	return -value
}

// IsPalindrome https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/train/go
// IsPalindrome Is it a palindrome?
func IsPalindrome(str string) bool {
	reversed := ""
	n := len(str) - 1
	for value := range str {
		reversed += strings.ToLower(string(str[n]))
		n--
		if n == 0 {
			reversed += strings.ToLower(string(str[0]))
			return strings.ToLower(str) == reversed
		}
		fmt.Println(value)
	}
	return strings.ToLower(str) == reversed
}

// PositiveSum https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go
// PositiveSum Sum of positive
func PositiveSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}

// NthEven https://www.codewars.com/kata/5933a1f8552bc2750a0000ed/train/go
// NthEven Get Nth Even Number
func NthEven(n int) int {
	return n*2 - 2
}

// RemoveChar https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go
// RemoveChar Remove First and Last Character
func RemoveChar(word string) string {
	return string(word[1 : len(word)-1])
}

// IsDivisible https://www.codewars.com/kata/5545f109004975ea66000086/train/go
// IsDivisible Is n divisible by x and y?
func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

// CorrectTail https://www.codewars.com/kata/56f695399400f5d9ef000af5/train/go
// CorrectTail Is this my tail?
func CorrectTail(body string, tail rune) bool {
	// if rune(body[len(body)-1]) == tail {
	// 	return true
	// }
	// return false

	return rune(body[len(body)-1]) == tail
}

// Goals https://www.codewars.com/kata/55f73be6e12baaa5900000d4/train/go
// Goals Grasshopper - Messi goals function
func Goals(goals ...int) int {
	sum := 0
	for _, goal := range goals {
		sum += goal
	}
	return sum
}

// func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
// 	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
// }

// ExpressionMatter https://www.codewars.com/kata/5ae62fcf252e66d44d00008e/train/go
// ExpressionMatter Expressions Matter
func ExpressionMatter(a int, b int, c int) int {
	// your code here
	solution := make([]int, 0)
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			solution = append(solution, a*(b+c))

		case 1:
			solution = append(solution, a*b*c)

		case 2:
			solution = append(solution, a+b*c)

		case 3:
			solution = append(solution, (a+b)*c)
		case 4:
			solution = append(solution, (a + b + c))
		}
	}
	return sort(solution)
}

func sort(n []int) int {
	max := 0
	for _, i := range n {
		if max < i {
			max = i
		}
	}

	return max
}

// MakeNegative https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go
// MakeNegative Return Negative
func MakeNegative(x int) int {
	return -int(math.Abs(float64(x)))
}

// MultipleOfIndex https://www.codewars.com/kata/5a34b80155519e1a00000009/train/go
// MultipleOfIndex Multiple of index
func MultipleOfIndex(ints []int) []int {
	// good luck
	res := []int{}
	for i, value := range ints {
		if i > 0 && value%i == 0 {
			res = append(res, value)
		}
	}
	return res
}

// Hero https://www.codewars.com/kata/59ca8246d751df55cc00014c/train/go
// Hero Is he gonna survive?
func Hero(bullets, dragons int) bool {
	return dragons*2 <= bullets
}

// HowManyDalmatians https://www.codewars.com/kata/56f6919a6b88de18ff000b36/train/go
// HowManyDalmatians 101 Dalmatians - squash the bugs, not the dogs!
func HowManyDalmatians(number int) string {
	dogs := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}
	if number <= 10 {
		return dogs[0]
	}
	if number <= 50 {
		return dogs[1]
	}
	if number < 101 {
		return dogs[2]
	}
	return dogs[3]
}

// Summation https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/go
// Summation Grasshopper - Summation
func Summation(n int) int {
	// the sleeper must awaken!
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

// Combat https://www.codewars.com/kata/586c1cf4b98de0399300001d/train/go
// Combat Grasshopper - Terminal game combat function
func Combat(health, damage float64) float64 {
	res := health - damage
	if res < 0 {
		return 0.0
	}
	return res
}

// BonusTime https://www.codewars.com/kata/56f6ad906b88de513f000d96/train/go
// BonusTime Do I get a bonus?
func BonusTime(salary int, bonus bool) string {
	if bonus {
		return "£" + strconv.Itoa(salary*10)
	}
	return "£" + strconv.Itoa(salary)
}

// Maps https://www.codewars.com/kata/57f781872e3d8ca2a000007e/train/go
// Maps Beginner - Lost Without a Map
func Maps(x []int) []int {
	res := []int{}
	for _, i := range x {
		res = append(res, i*2)
	}
	return res
}

// OddCount https://www.codewars.com/kata/59342039eb450e39970000a6/train/go
// OddCount Count Odd Numbers below n
func OddCount(n int) int {
	if n%2 == 1/2 {
		return n / 2
	}
	return n/2 + 1/2
}

// Xor https://www.codewars.com/kata/56fa3c5ce4d45d2a52001b3c/train/go
// Xor Exclusive "or" (xor) Logical Operator
func Xor(a, b bool) bool {
	// if a && !b {
	// 	return true
	// }
	// if !a && !b {
	// 	return false
	// }
	// if !a && b {
	// 	return true
	// }
	// return false
	if a != b {
		return true
	}
	return false
}

// Multiply https://www.codewars.com/kata/523b66342d0c301ae400003b/train/go
// Multiply Function 3 - multiplying two numbers
func Multiply(x, y int) int {
	return x * y
}

// SquareOrSquareRoot https://www.codewars.com/kata/57f6ad55cca6e045d2000627/train/go
// SquareOrSquareRoot To square(root) or not to square(root)
func SquareOrSquareRoot(arr []int) []int {
	transformed := []int{}
	for _, i := range arr {
		root := math.Sqrt(float64(i))
		if (math.Ceil(root) * root) == float64(i) {
			transformed = append(transformed, int(float64(root)))
		} else {
			transformed = append(transformed, int(float64(i*i)))
		}
	}
	return transformed
}

// GetGrade https://www.codewars.com/kata/55cbd4ba903825f7970000f5/train/go
// GetGrade Grasshopper - Grade book
func GetGrade(a, b, c int) rune {

	res := 'F'
	switch {
	case mean(a, b, c) >= 90:
		res = 'A'
	case mean(a, b, c) >= 80:
		res = ('B')
	case mean(a, b, c) >= 70:
		res = ('C')
	case mean(a, b, c) >= 60:
		res = ('D')
	case mean(a, b, c) >= 0:
		res = ('F')
	}
	return res
}

func mean(a, b, c int) int {
	return (a + b + c) / 3
}

// Greet https://www.codewars.com/kata/523b4ff7adca849afe000035/train/go
// Greet Function 1 - hello world
func Greet() string {
	return "hello world!"
}

// SeatsInTheater https://www.codewars.com/kata/588417e576933b0ec9000045/train/go
// SeatsInTheater Simple Fun #1: Seats in Theater
func SeatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}

// MyString type
type MyString string

// IsUpperCase https://www.codewars.com/kata/56cd44e1aa4ac7879200010b/train/go
// IsUpperCase Is the string uppercase?
func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))

}

// CountSheep https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go
// CountSheep If you can't sleep, just count sheep!!
func CountSheep(num int) string {
	counts := ""
	for i := 1; i <= num; i++ {
		counts += strconv.Itoa(i) + " sheep..."
	}
	return counts
}

// Grow https://www.codewars.com/kata/57f780909f7e8e3183000078/train/go
// Grow Beginner - Reduce but Grow
func Grow(arr []int) int {
	sum := 1
	for _, number := range arr {
		sum = sum * number
	}
	return sum
}

// OtherAngle https://www.codewars.com/kata/5a023c426975981341000014/train/go
// OtherAngle Third Angle of a Triangle
func OtherAngle(a int, b int) int {
	return 180 - (a + b)
}

// Move https://www.codewars.com/kata/563a631f7cbbc236cf0000c2/train/go
// Move Grasshopper - Terminal game move function
func Move(position int, roll int) int {
	return position + roll*2
}

// MakeUpperCase https://www.codewars.com/kata/57a0556c7cb1f31ab3000ad7/train/go
// MakeUpperCase func
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}

// CheckForFactor https://www.codewars.com/kata/55cbc3586671f6aa070000fb/train/go
// CheckForFactor Grasshopper - Check for factor
func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}

// MonkeyCount https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/go
// MonkeyCount Count the Monkeys!
func MonkeyCount(n int) []int {
	monkeys := []int{}
	for i := 1; i <= n; i++ {
		monkeys = append(monkeys, i)
	}
	return monkeys

}

// Century https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/train/go
// Century Century From Year
func Century(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return year/100 + 1
}

// Solution https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
// Solution Reversed Strings
func Solution(word string) string {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}
