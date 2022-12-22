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

  describe "#has_same_coordinates_with" do
    context "when two GridPoint object have same coordinates" do
      it "should return true" do
        expect(obj.has_same_coordinates_with(GridPoint.new(x, y))).to eq true
      end
    end

    context "when two GridPoint object have different coordinates" do
      it "should return true" do
        expect(obj.has_same_coordinates_with(GridPoint.new(x+1, y+1))).to eq false
      end
    end
  end
end
