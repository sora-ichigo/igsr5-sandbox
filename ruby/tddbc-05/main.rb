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

  def neighbor_of?(grid_point)
    condition1 = self.x-1 == grid_point.x && self.y == grid_point.y
    condition2 = self.x+1 == grid_point.x && self.y == grid_point.y
    condition3 = self.x == grid_point.x && self.y-1 == grid_point.y
    condition4 = self.x == grid_point.x && self.y+1 == grid_point.y

    condition1 || condition2 || condition3 || condition4
  end
end

class GridPointSet
  attr_accessor :grid_points

  # @param [Array<GridPoint>] grid_points
  def initialize(*grid_points)
    @grid_points = grid_points
  end

  # @param [GridPoint] target_grid_point
  def contain?(target_grid_point)
    self.grid_points.reduce(false) do |is_contain, grid_point|
      is_contain || grid_point.same_coordinates_with?(target_grid_point)
    end
  end

  def connected?
    connected = true
    self.grid_points.each_cons(2) do |grid_points|
      connected = connected && grid_points[0].neighbor_of?(grid_points[1])
    end

    connected
  end

  def count
    self.grid_points.count
  end

  def traversable?

  end
end
