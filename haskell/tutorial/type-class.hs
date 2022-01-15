inc :: Float -> Float
inc x = x + 1

-- 型クラス: Num, ...
-- (これ自体は方ではなく型の集合)
-- (関数やインスタンスを定義できる)

-- 型、型クラスのインスタンス
-- Integer, Float
-- (実際の変数などに割当てることができる型)

-- NOTE: (+) には型クラス制約 Num があり、型変数 a がそれを満たさない
-- e.g.
-- add2 :: a -> a -> a
-- add2 i j = i + j

main = do
print $ inc 5
