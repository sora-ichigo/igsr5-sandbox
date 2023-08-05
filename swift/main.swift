let hoge = "80"
print(hoge, "hello")

var str: String

str = "3"
print(str)

var num = 3
let fol = 3.14

let a = 3 + 3.14

print(Int8.max)

typealias MyInt = Int

var myInt: MyInt = 0
print(myInt)

let b = true

if b {
    print("true")
} else {
    print("false")
}

let http404error = (404, "Not Found")

let (statusCode, statusMessage) = http404error
print(statusCode)
print(statusMessage)

let possibleNumber = "123"
let convertedNumber = Int(possibleNumber)

if convertedNumber != nil {
    print(convertedNumber!)
}

func printNumber(_ numberString: String?) {
    guard let number = numberString else {
        print("入力はnilまたは無効です")
        return
    }
    print("数字は: \(number)")
}

printNumber("3")
printNumber(nil)

let numbers = [1, 2, 3, 4, 5]
print("numbers: \(numbers)")

for number in numbers {
    print(number)
}
print("")


for number in numbers[2...] {
    print(number)
}

print("")
for i in 0...5 {
    print(i)
}

print("")

for i in 0..<5 {
    print(i)
}
