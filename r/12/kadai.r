# filter 関数を利用する
# `install.packages("dplyr")` は済み
# library(dplyr)
# 
# # 各品種のデータを取得
# setoa <- filter(iris, grepl("setosa", Species))
# versicolor <- filter(iris, grepl("versicolor", Species))
# virginica <-  filter(iris, grepl("virginica", Species))
# 
# # F検定
# # 有意水準 α = 0.05
# 
# var.test(setoa$Petal.Width, versicolor$Petal.Width, var.equal=FALSE)
# var.test(setoa$Petal.Width, virginica$Petal.Width, var.equal=FALSE)
# var.test(versicolor$Petal.Width, virginica$Petal.Width, var.equal=FALSE)

s = iris$Species
w = iris$Sepal.Width
TukeyHSD(aov(w~s))
