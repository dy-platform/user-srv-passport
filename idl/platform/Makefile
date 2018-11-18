.PHONY:all clean echo

dirs := $(shell ls -l | grep '^d' |awk '{print $$9}')

objects := $(patsubst %.proto, %.pb.go, $(wildcard *.proto))


all: $(objects) $(dirs)

$(dirs):echo
	make -C $@

echo:
	@echo

%.pb.go: %.proto
	protoc -I. --go_out=. $^

clean:
	rm -f