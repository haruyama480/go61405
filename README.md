# go61405
inspired: https://github.com/nobishino/go61405

# Study new range over proposal

Sample codes for studying the proposal https://github.com/golang/go/issues/61405. 

## Install gotip and build with CL

```
go install golang.org/dl/gotip@latest
gotip download 510541   # download and build CL 510541              
gotip version  # should say "(w1/ rangefunc)"  
```

## Execute sample code

```
gotip run .
```