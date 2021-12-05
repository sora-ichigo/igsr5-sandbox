# -----------------------------
# define_method
# -----------------------------
class Computer
  def initialize(name_list)
    @name_list = name_list
    @name_list.each do |name|
      Computer.define_component name
    end
  end

  def self.define_component(name)
    define_method(name) do
      p name
    end
  end
end

computer = Computer.new(%w[mouse keyboard])
computer.mouse

# -----------------------------
# missing_method
# -----------------------------

class Lawyer
  def method_missing(method, *args)
    p method.to_s
  end

  def respond_to_missing?(method, include_private = false)
    true
  end
end

lawyer = Lawyer.new
lawyer.keyboard
p lawyer.respond_to?(:keyboard)
