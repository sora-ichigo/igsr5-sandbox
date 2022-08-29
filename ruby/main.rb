puts "お腹すいた"
puts "なんか食べたい"

a =  [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
b =  [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
c =  [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
def main
  a.each do |i|
    puts i
  end

  b.each do |i|
    puts i*2
  end

  print_c
end

def print_c
  c.each do |i|
    puts i*3
  end
end
