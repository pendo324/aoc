EXT:=""
RM:=""

ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del 
else
	RM=rm
endif

.PHONY: exe
exe:
	go build -o aoc$(EXT)

all: exe

.PHONY: clean
clean:
	$(RM) aoc$(EXT)
