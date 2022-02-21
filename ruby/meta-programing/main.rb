class String
  def count(n)
    0
  end

  def original_method
    "original"
  end
end

a = "hello"

p a.count "l"

p a.original_method
