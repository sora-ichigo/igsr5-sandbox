require_relative "../main.rb"

describe "GridPoint" do
  let(:x) { 4 }
  let(:y) { 7 }
  let(:obj) { GridPoint.new(x, y) }

  describe "#x" do
    it { expect(obj.x).to eq x }
  end

  describe "#y" do
    it { expect(obj.y).to eq y }
  end

  describe "#notation" do
    it "should return '(x,y)' as String" do
      expect(obj.notation).to eq "(#{x},#{y})"
    end
  end

  describe "#same_coordinates_with?" do
    context "when two GridPoint object have same coordinates" do
      it "should return true" do
        expect(obj.same_coordinates_with?(GridPoint.new(x, y))).to eq true
      end
    end

    context "when two GridPoint object have different coordinates" do
      it "should return true" do
        expect(obj.same_coordinates_with?(GridPoint.new(x+1, y+1))).to eq false
      end
    end
  end

  # 隣り合う条件 (x-1,y),(x+1,y),(x,y-1),(x,y+1)
  describe "#neighbor_of?" do
    it "(x-1,y) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x-1, y))).to eq true
    end
    it "(x+1,y) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x+1, y))).to eq true
    end
    it "(x,y-1) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x, y-1))).to eq true
    end
    it "(x,y+1) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x, y+1))).to eq true
    end
    it "(x+1,y+1) is not neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x+1, y+1))).to eq false
    end
  end
end

describe "GridPointSet" do
  context "when grid_point_set has two grid_points" do
    let(:grid_point1) { GridPoint.new(1, 2) }
    let(:grid_point2) { GridPoint.new(3, 6) }
    let(:grid_point_set) { GridPointSet.new(grid_point1, grid_point2) }

    it do
      expect(grid_point_set.grid_points).to eq [grid_point1, grid_point2]
    end

    describe "#contain?" do
      context "when contains given grid_points contain" do
        it "should return true" do
          expect(grid_point_set.contain?(grid_point1)).to eq true
        end
      end
      context "when doesn't contains given grid_points contain" do
        it "should return false" do
          expect(grid_point_set.contain?(GridPoint.new(10, 20))).to eq false
        end
      end
    end

    # 隣り合う条件 (x-1,y),(x+1,y),(x,y-1),(x,y+1)
    describe "#connected?" do
      subject { grid_point_set.connected? }
      context "when (x-1,y) pattern" do
        let(:grid_point1) { GridPoint.new(1, 2) }
        let(:grid_point2) { GridPoint.new(0, 2) }
        it "should return true" do
          is_expected.to eq true
        end
      end
      context "when (x+1,y) pattern" do
        let(:grid_point1) { GridPoint.new(1, 2) }
        let(:grid_point2) { GridPoint.new(2, 2) }
        it "should return true" do
          is_expected.to eq true
        end
      end
      context "when (x,y-1) pattern" do
        let(:grid_point1) { GridPoint.new(1, 2) }
        let(:grid_point2) { GridPoint.new(1, 1) }
        it "should return true" do
          is_expected.to eq true
        end
      end
      context "when (x,y+1) pattern" do
        let(:grid_point1) { GridPoint.new(1, 2) }
        let(:grid_point2) { GridPoint.new(1, 3) }
        it "should return true" do
          is_expected.to eq true
        end
      end
      context "when grid_points doesn't connect" do
        let(:grid_point1) { GridPoint.new(1, 2) }
        let(:grid_point2) { GridPoint.new(10, 20) }
        it "should return false" do
          is_expected.to eq false
        end
      end
    end
  end

  context "when grid_point_set has three grid_points" do
    let(:grid_point1) { GridPoint.new(1, 2) }
    let(:grid_point2) { GridPoint.new(3, 6) }
    let(:grid_point3) { GridPoint.new(9, 18) }
    let(:grid_point_set) { GridPointSet.new(grid_point1, grid_point2, grid_point3) }

    it do
      expect(grid_point_set.grid_points).to eq [grid_point1, grid_point2, grid_point3]
    end

    describe "#contain?" do
      context "when contains given grid_points contain" do
        it "should return true" do
          expect(grid_point_set.contain?(grid_point3)).to eq true
        end
      end
      context "when doesn't contains given grid_points contain" do
        it "should return false" do
          expect(grid_point_set.contain?(GridPoint.new(10, 20))).to eq false
        end
      end
    end

    # 格子点集合に含まれる任意の格子点について、
    # その格子点から隣り合っている格子点のみを経由して
    # その他全ての格子点へ到達できる場合に限り、
    # その格子点集合が連結しているものとします
    describe "#connected?" do
      context "when grid_point_set is connected" do
        context "ver1" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(2, 2) }
          let(:grid_point3) { GridPoint.new(3, 2) }

          it "should return true" do
            expect(grid_point_set.connected?).to eq true
          end
        end

        context "ver1" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(1, 3) }
          let(:grid_point3) { GridPoint.new(1, 2) }

          it "should return true" do
            expect(grid_point_set.connected?).to eq true
          end
        end

        context "ver3" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(1, 3) }
          let(:grid_point3) { GridPoint.new(2, 3) }

          it "should return true" do
            expect(grid_point_set.connected?).to eq true
          end
        end
      end

      context "when grid_point_set is not connected" do
        context "ver1" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(2, 2) }
          let(:grid_point3) { GridPoint.new(4, 2) }

          it "should return false" do
            expect(grid_point_set.connected?).to eq false
          end
        end

        context "ver2" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(3, 2) }
          let(:grid_point3) { GridPoint.new(4, 2) }

          it "should return false" do
            expect(grid_point_set.connected?).to eq false
          end
        end

        context "ver3" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(1, 3) }
          let(:grid_point3) { GridPoint.new(4, 2) }

          it "should return false" do
            expect(grid_point_set.connected?).to eq false
          end
        end
      end
    end

    describe "#count" do
      it "should return grid_points count" do
        expect(grid_point_set.count).to eq 3
      end
    end
  end

  context "when grid_point_set has four grid_points" do
    let(:grid_point1) { GridPoint.new(1, 2) }
    let(:grid_point2) { GridPoint.new(3, 6) }
    let(:grid_point3) { GridPoint.new(9, 18) }
    let(:grid_point4) { GridPoint.new(27, 54) }
    let(:grid_point_set) { GridPointSet.new(grid_point1, grid_point2, grid_point3, grid_point4) }

    it do
      expect(grid_point_set.grid_points).to eq [grid_point1, grid_point2, grid_point3, grid_point4]
    end

    describe "#contain?" do
      context "when contains given grid_points contain" do
        it "should return true" do
          expect(grid_point_set.contain?(grid_point3)).to eq true
        end
      end
      context "when doesn't contains given grid_points contain" do
        it "should return false" do
          expect(grid_point_set.contain?(GridPoint.new(10, 20))).to eq false
        end
      end
    end

    # 格子点集合に含まれる任意の格子点について、
    # その格子点から隣り合っている格子点のみを経由して
    # その他全ての格子点へ到達できる場合に限り、
    # その格子点集合が連結しているものとします
    describe "#connected?" do
      context "when grid_point_set is connected" do
        context "ver1" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(2, 2) }
          let(:grid_point3) { GridPoint.new(3, 2) }
          let(:grid_point4) { GridPoint.new(4, 2) }

          it "should return true" do
            expect(grid_point_set.connected?).to eq true
          end
        end

        context "ver2" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(1, 3) }
          let(:grid_point3) { GridPoint.new(1, 2) }
          let(:grid_point4) { GridPoint.new(2, 2) }

          it "should return true" do
            expect(grid_point_set.connected?).to eq true
          end
        end
      end

      context "when grid_point_set is not connected" do
        context "ver1" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(2, 2) }
          let(:grid_point3) { GridPoint.new(4, 2) }
          let(:grid_point4) { GridPoint.new(5, 2) }

          it "should return false" do
            expect(grid_point_set.connected?).to eq false
          end
        end

        context "ver2" do
          let(:grid_point1) { GridPoint.new(1, 2) }
          let(:grid_point2) { GridPoint.new(2, 2) }
          let(:grid_point3) { GridPoint.new(3, 2) }
          let(:grid_point4) { GridPoint.new(5, 2) }

          it "should return false" do
            expect(grid_point_set.connected?).to eq false
          end
        end
      end
    end

    describe "#count" do
      it "should return grid_points count" do
        expect(grid_point_set.count).to eq 4
      end
    end
  end
end
