.PHONY:all

objects := $(patsubst %.proto, %.pb.go, $(wildcard *.proto))

all:  $(objects)


%.pb.go:%.proto
	protoc -I. --go_out=. --micro_out=. $^