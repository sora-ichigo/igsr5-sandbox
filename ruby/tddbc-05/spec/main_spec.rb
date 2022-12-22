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
  end
end
