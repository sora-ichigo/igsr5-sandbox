main = print "This is example Monado"

-- モナドとは
-- 型クラス Monad に属すること
-- 下記を満たせば良い
-- 1. リストのように型引数を一つ受け取って作られるクラス
-- 2. >>= が定義されている（バインド関数）
-- 3. return が定義されている（データコンストラクタ）

-- ref. 
-- *Main> :i Monad
-- type Monad :: (* -> *) -> Constraint
-- class Applicative m => Monad m where
--   (>>=) :: m a -> (a -> m b) -> m b
--   (>>) :: m a -> m b -> m b
--   return :: a -> m a
--   {-# MINIMAL (>>=) #-}
--         -- Defined in ‘GHC.Base’
-- instance Monad (Either e) -- Defined in ‘Data.Either’
-- instance Monad [] -- Defined in ‘GHC.Base’
-- instance Monad Maybe -- Defined in ‘GHC.Base’
-- instance Monad IO -- Defined in ‘GHC.Base’
-- instance Monad ((->) r) -- Defined in ‘GHC.Base’
-- instance (Monoid a, Monoid b, Monoid c) => Monad ((,,,) a b c)
--   -- Defined in ‘GHC.Base’
-- instance (Monoid a, Monoid b) => Monad ((,,) a b)
--   -- Defined in ‘GHC.Base’
-- instance Monoid a => Monad ((,) a) -- Defined in ‘GHC.Base’

-- [>>=]
-- バインド関数
-- *Main> :type (>>=)
-- (>>=) :: Monad m => m a -> (a -> m b) -> m b
-- ↑ m は型コンストラクタ（型を引数として受け取って型を作る。e.g. []）

list = [1, 2, 3, 4, 5]
makeDoublePair = \x -> [x, x*2]
-- *Main> list >>= makeDoublePair
-- [1,2,2,4,3,6,4,8,5,10]

-- do 構文は >>= のネストを緩和するための糖衣構文
ioTest = do
    firstName <- getLine
    lastName <- getLine
    putStrLn ("Hello, " ++ firstName ++ lastName)
