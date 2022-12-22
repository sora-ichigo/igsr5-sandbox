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
end
