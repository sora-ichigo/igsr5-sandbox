n=1000
m<-numeric(length=10000)
for(i in 1:10000){
  h<-rnorm(n, mean=0, sd=1)
  m[i]<- mean(h)
}

hist(m)
