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
-- as パターン。引数の後に@ をつけることでパターンマッチの対象の値を参照できる
firstLetter :: String -> String
firstLetter "" = "Emptry string."
firstLetter all@(x:xs) = "The first letter of " ++ all ++ " is " ++ [x]

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

-- [where]
-- 定義した*関数*内に変数を束縛できる
-- where キーワードを使わない場合
totalCheck :: Integer -> Integer -> String
totalCheck x y
    | x + y < 10 = "less than 10"
    | x + y < 100 = "less than 100"
    | otherwise = "a large number"

-- where キーワードを使う場合
totalCheck' :: Integer -> Integer -> String
totalCheck' x y
    | total < 10 = "less than 10"
    | total < 100 = "less than 100"
    | otherwise = "a large number"
    where total = x + y

-- [let]
-- どこでも変数を束縛できる
-- where は関数の終わりのみ

-- [ラムダ式]
-- 無名関数的なやつ
addThree :: Int -> Int -> Int -> Int
addThree x y z = x + y + z
-- ↑ と同じ。引数一つしか無理なのでカリー化してる。
addThree' = \x -> \y -> \z -> x+y+z

-- [畳み込み]
-- foldl: 左畳み込み
-- foldr: 右畳み込み
foldlEg = foldl (+) 0 [1, 2, 3]
-- 6

-- [$]
-- 右結合な関数適用
-- sum map sqrt [1, 2, 3] とは書けない。関数は左から評価されていくため
test_right_merge = sum $ map sqrt [1, 2 ,3]

-- [関数合成]
-- . で関数を連結することで数学みたいな関数合成できる
f = \x -> x*x
g = \x -> x+1
h = (f.g) 1
-- f(g(1))
-- 4
