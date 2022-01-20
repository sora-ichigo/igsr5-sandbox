def url_of_images(*args)
  p args
end

url_of_images("640x")
# → ["640x"]
url_of_images("640x300#")
# → ["640x300#"]

def url_of_images_1(w: nil, h: nil)
  p w, h
end
url_of_images_1(**{w: 100, h:200})

