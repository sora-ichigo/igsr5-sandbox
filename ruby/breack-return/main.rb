require 'securerandom'

def generate_url_key_1
  loop do
    url_key = 'foo'
    break url_key
  end
end

def generate_url_key_2
  url_key = ''
  loop do
    url_key = 'boo'
    break
  end
  url_key
end

p generate_url_key_1 # → foo
p generate_url_key_2 # → boo
