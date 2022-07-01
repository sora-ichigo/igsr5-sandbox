require 'rubygems'
require 'json'
require 'bundler/setup'

client = HTTPClient.new

APIKEY = ENV.fetch("APIKEY")
base_url = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

# 豊田市
latitude = "35.109783"
longitude = "137.166782"

place_params = {
  key: APIKEY,
  location: "#{latitude},#{longitude}",
  radius: 10000,
  keyword: "温泉",
  language: 'ja'
}

res = client.get(base_url, place_params)

JSON.parse(res.body)["results"].map do |r|
  puts "name: #{r["name"]}"
  puts "types: #{r["types"]}"
  puts "vicinity: #{r["vicinity"]}"
  puts "---------------------------------------"
end