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
Test.hello_c
# output
# chap4.rb:18:in `hello_c': undefined local variable or method `c' for Test:Class (NameError)
