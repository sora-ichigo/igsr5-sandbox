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
end
