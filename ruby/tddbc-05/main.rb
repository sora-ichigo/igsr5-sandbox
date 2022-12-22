# from. http://devtesting.jp/tddbc/?TDDBC%E4%BB%99%E5%8F%B005%2F%E8%AA%B2%E9%A1%8C

class GridPoint
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def notation
    "(#{x},#{y})"
  end

  def same_coordinates_with?(grid_point)
    self.x == grid_point.x && self.y == grid_point.y
  end
end
