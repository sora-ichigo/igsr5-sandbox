lData <- read.csv("./laureates.csv", header=T, fileEncoding = "cp932")
attach(lData)
# NOTE: available data
# Chocolate
# Laureates

r<-lm(Laureates~Chocolate)

plot(Chocolate, Laureates)
abline(r)


