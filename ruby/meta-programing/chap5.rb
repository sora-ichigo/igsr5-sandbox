class MyClass
  puts 'Hello'
end

class C
  def m1
    def m2; end
  end
end

class D < C; end

obj = D.new
p obj.m1

p C.instance_methods(false)

def add_method_to(a_class)
  a_class.class_eval do
    def m; 'Hello!'; end
  end
end

add_method_to String
p "abc".m

class MyClass2
  @my_var = 1
  def self.read; @my_var; end
  def write; @my_var = 2; end
  def read; @my_var; end
end

obj = MyClass2.new
p obj.read
p obj.write
p obj.read
p MyClass2.read

@@v = 1

class MyClass3
  @@v = 2
end

p @@v

# クイズ、これを class 使わずにかけ
# class MyClass < Array
#   def my_method
#     'Hello!'
#   end
# end

# a = Array
# a.class_eval { def self.my_method; 'Hello!'; end }
# 
# p Array.my_method

c = Class.new(Array) do
  def my_method
    "Hello!"
  end
end

MyClass = c

d = MyClass.new
p d.my_method # ← "Hello!"
p c.name # ← MyClass

a = "abc"
def a.hello
  p "hello"
end

a.hello
# "abc".hello
# cause error: undefined method `hello`

class Book
  def title
    p "title"
  end

  def self.deprecate(old_method, new_method)
    define_method(old_method) do |*args, &block|
      warn "Warning: #{old_method}() is deprecated. Use #{new_method}()."
      send(new_method, *args, &block)
    end
  end

  deprecate :GetTitle, :title
end

book = Book.new
book.title # → "title"
book.GetTitle # → "title"

class MyClass4
  p self # → MyClass4
  def method
    p self # → <MyClass:0x00000001230fbda0>
    # self.singleton_method_test はエラーになる
    self.class.singleton_method_test # → "hello"
  end

  def self.singleton_method_test
    p "hello"
  end
end

m = MyClass4.new
m.method

class C
  def a_method
    'C#a_method()'
  end
end

class D < C
end

obj = D.new
p obj.a_method
p obj.singleton_class

class << obj
  def a_singleton_method
    'obj#a_singleton_method()'
  end

  def b_singleton_method
    'obj#b_singleton_method()'
  end
end

class << obj
  def c_singleton_method
    'obj#c_singleton_method()'
  end
end

p obj.a_singleton_method
p obj.b_singleton_method
p obj.c_singleton_method

p obj.singleton_class.superclass
p D.singleton_class
p D.singleton_class.superclass

module MyModule
  def my_method; "hello";end
end

class E
  include MyModule
  class << self
    include MyModule
  end
end

class F
  extend MyModule
end


p E.instance_methods
p E.singleton_methods

p F.instance_methods
p F.singleton_methods

# クイズ 1 + 1, 2 + 3 などどんな足し算をしても正しい結果 + 1 されるようにしろ
# ヒント 1 + 1 は実際には Fixnum#+ を呼び出している
# e.g.  1.+(1)
p 1.class
p 1.+(1)

p 10000.is_a?(Fixnum)

class Fixnum
  alias_method :tmp, :+
  def +(a)
    self.tmp(a).tmp(1)
  end
end

# できなかった
# module FixnumExtend
#   def +(value)
#     a = super(value)
#   end
# end
# class Fixnum
#   prepend FixnumExtend
# end
#
# できなかった
# module FixnumRefinements
#   refine Fixnum do
#     def +(value)
#      result = super(value)
#      super(1)
#     end
#   end
# end

using FixnumRefinements
p 1000 + 1000
