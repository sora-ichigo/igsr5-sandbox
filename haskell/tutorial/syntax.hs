main = print "this is Haskell's syntax example."

-- [if 文]
-- Haskell における if は式扱い
-- なので必ず値を返す (= `else` が必須)

fanc x = if x > 5 then x else 5

-- [List]
-- 同じ型の要素を格納できる。e.g. [Int]

integers = [1, 2, 3, 4, 5, 6]
-- *Main> integers
-- [1,2,3,4,5,6]
-- *Main> :t integers
-- integers :: [Integer]

-- NOTE: リスト内包表記
a = [2*x | x <- [1,2,3,4,5]]
-- *Main> a
-- [2,4,6,8,10]

b = [2*x | x <- [1,2,3,4,5], x < 3]
-- 内包表記の条件は複数指定できる。これをフィルタするという。
-- *Main> b
-- [2,4]

-- [タプル]
-- NOTE: 複数の違う方を格納できることをヘテロであるという
c = (1, "String", 'c', 3.14)

-- [パターンマッチ]
-- データ型が従うべきパターンを指定してそのパターンに従ってデータを分解するために従う
factorial :: Int -> Int
factorial 0 = 1
factorial n = n * factorial (n-1)

-- [ガード]
-- 他の言語でいう Switch 文みたいなやつ

fact :: Integer -> Integer
fact n
    | n < 0 = error "error!!"
    | n == 0 = 1
    | otherwise = n * fact(n - 1)

bmi :: Float -> String
bmi n
    | n <= 18.5 = "smart"
    | n <= 25.0 = "nomarl"
    | n > 25.0 = "yabai"
