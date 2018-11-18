.PHONY:all clean
exclude_dirs := Makefile
dirs := $(shell ls)
dirs := $(filter-out $(exclude_dirs),$(dirs))
all:
	$(foreach N,$(dirs),make -C $(N);)