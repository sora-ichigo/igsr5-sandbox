main = print "This is example Monado"

-- モナドとは
-- 1. リストのように型引数を一つ受け取って作られる型クラス

-- [>>=]
-- バインド関数
list = [1, 2, 3, 4, 5]
makeDoublePair = \x -> [x, x*2]
-- *Main> list >>= makeDoublePair
-- [1,2,2,4,3,6,4,8,5,10]
