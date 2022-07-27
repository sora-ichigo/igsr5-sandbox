# 1====================
# txtData <- read.csv("./0721_TSH.csv", header=T, fileEncoding = "cp932")
# attach(txtData)
# v<-txtData$TSH
# 
# mean(v) # 平均値
# median(v) # 中央値
# sd(v) # 標準偏差
# 
# a_v <- v[v < 24.30] # 外れ値を消す
# mean(a_v) # 平均値
# median(a_v) # 中央値
# 
# sd(a_v) # 標準偏差

# 2====================
# x<-runif(20, 0, 3)
# e<-rnorm(20, mean=0, sd=1)
# x
# e
# y<-x+e
# plot(x, y)
# r<-lm(y~x)
# abline(r)
# summary(r)
# cor(x, y)

# 3====================
m<-1.36+(5/200)
length<-iris$Petal.Length

qt(0.025, 49)
qt(0.975, 49)
t.test(length[1:50], mu=m, var.equal=TRUE)
