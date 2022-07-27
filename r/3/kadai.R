txtData <- read.csv("./3/kadai3.csv", header=T, fileEncoding = "cp932")
attach(txtData)

english <- txtData$英語
japanese <- txtData$国語
math <- txtData$数学
science <- txtData$理科
society <- txtData$社会

# 1. 平均値
mean(science)

# 2. 不変分散
var(english)
#sum((english-mean(english))^2)/(length(english)-1)

# 3. 標準偏差(普遍分散の平方根)
sd(math)

# 4. 各中央値の最大値
max(c(median(english), median(japanese), median(math), median(science), median(society)))

# 5. 平均値と中央値の差が最小の科目名
print("5. 各科目の平均値と中央値")
print("国語")
mean(japanese)-median(japanese)
print("英語")
mean(english)-median(english)
print("数学")
mean(math)-median(math)
print("理科")
mean(science)-median(science)
print("社会")
mean(society)-median(society)

# 6.
magic_number <- 4 + 22 - 2

# z得点 = (得点 - 平均値) / 標準偏差
# 平均値が0になって標準偏差が1になる
sc_z <- (science - mean(science)) / sd(science)
# 偏差値
# 平均が 50, 標準偏差が 10 になるように調整したもの
(sc_z * 10 + 50)[magic_number]

# 7.
magic_number <- 4 + 22 - 1
e_z <- (english - mean(english)) / sd(english)
(e_z * 10 + 50)[magic_number]

# 8.
magic_number <- 4 + 22
m_z <- (math - mean(math )) / sd(math)
(m_z * 10 + 50)[magic_number]

# 9.
magic_number <- 4 + 22 + 1
j_z <- (japanese - mean(japanese)) / sd(japanese)
(j_z * 10 + 50)[magic_number]

# 10.
magic_number <- 4 + 22 + 2
so_z <- (society - mean(society)) / sd(society)
(so_z * 10 + 50)[magic_number]
