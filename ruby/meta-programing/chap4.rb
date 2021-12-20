a = "a"
@b = "b"
c = "c"

def hello
  p "hello"
end

class Test
  a = "a_test"
  @b = "b_test"

  def hello
    p "hello_test"
  end

  def self.hello_c
    p c
  end

  def hello_b
    p @b
  end

  def self.method
    hello
    yield
  end

  def self.method2
    hello
    yield
  end

  def method2
    hello_b
    hello
    yield
  end
end

p "1"
Test.method { p a }
# output
# "hello"
# "a"

p "2"
Test.method2 do
  hello
  p @b
end
# output
# "hello"
# "hello"
# "b"

p "3"
test = Test.new
test.method2 do
  hello
  p @b
end
# output
# "hello_test"
# "hello"
# "b"

p "4"
#Test.hello_c
# output
# chap4.rb:18:in `hello_c': undefined local variable or method `c' for Test:Class (NameError)

class MyClass
  def initialize
    @v = 1
  end

  def test
    yield
  end
end

@obj = MyClass.new

@obj.instance_eval do
  p self
  p @v
end
# output
# MyClass::---
# 1

@obj.test do
  p self
  p @v
end
# output
# main
# nil

inc = Proc.new { |x| x+ 1 }
p inc.call(2)
# output
# 3

inc = lambda { |x| x + 2 }
p inc.call(2)
# output
# 4

inc = ->(x){ x + 2 }
p inc.call(3)
# output
# 5

def math(x, y)
  yield(x, y)
end

def do_math(x, y, &operation)
  math(x, y, &operation)
end

p do_math(2, 3) {|x, y| x * y}
# output
# 6

def test
  yield
end

def test_1
  p "bb"
end

test { p "aa" }
:test_1.to_proc.call
