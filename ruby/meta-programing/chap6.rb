# user_input = "User input: #{gets()}"
# p user_input
# a = gets()
# puts user_input.tainted?
# puts a.tainted?

# add_checked_attribute
# 1

# def add_checked_attribute(klass, attribute)
#   eval "
#     class #{klass}
#       def #{attribute}=(value)
#         raise 'Invalid attribute' unless value
#         @#{attribute} = value
#       end
# 
#       def #{attribute}
#         @#{attribute}
#       end
#     end
#   "
# end

# 2

def add_checked_attribute(klass, attribute, &validation)
  klass.class_eval do
    define_method "#{attribute}" do
      instance_variable_get "@#{attribute}"
    end

    define_method "#{attribute}=" do |value|
      raise "Invalid attribute Error: #{value}" unless validation.call(value)
      instance_variable_set("@#{attribute}", value)
    end
  end
end

add_checked_attribute(String, :my_attr) { |v| v > 18 }
a = String.new

p a.my_attr
# a.my_attr= 11 → Invalid attribute Error
a.my_attr = 20 # → OK
p a.my_attr
# a.my_attr = nil → Invalid attribute Error

# 4

# class Class
#   def attr_checked(attribute, &validation)
#     define_method "#{attribute}" do
#       instance_variable_get "@#{attribute}"
#     end
# 
#     define_method "#{attribute}=" do |value|
#       raise "Invalid attribute Error: #{value}" unless validation.call(value)
#       instance_variable_set("@#{attribute}", value)
#     end
#   end
# end

# 5

module CheckedAttributes
  def self.included(base)
    base.extend ClassMethods
  end

  module ClassMethods
    def attr_checked(attribute, &validation)
      define_method "#{attribute}" do
        instance_variable_get "@#{attribute}"
      end

      define_method "#{attribute}=" do |value|
        raise "Invalid attribute Error: #{value}" unless validation.call(value)
        instance_variable_set("@#{attribute}", value)
      end
    end
  end
end

class A
  include CheckedAttributes
  # extend ClassMethods ← 普通に module を用意して extend でも動くよね

  attr_checked :my_attr do |v|
    v > 20
  end
end

a = A.new
p a.my_attr
a.my_attr = 21
p a.my_attr

# 5
