package calculator

var logMessage = "[LOG]"

// Version
var Version = "0.2.0"

func internalSum(number int) int {
  return number - 1
}

func Sum(number1 int, number2 int) int {
  return number1 + number2
}
