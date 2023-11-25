default: help

year=2015
day?=40

DIRECTORY:=day$(day)

part?=1
answer?=0


help:
	@echo "\n\n*****************"
	@echo "* Download a new puzzle file and create a directory"
	@echo "* with a rust project"
	@echo ""
	@echo "make newday day=01 year=2015"

	@echo "\n*****************"

	@echo "* Submit a solution for a given day/year"
	@echo ""
	@echo "make submit day=01 part=1 answer=1234"

	@echo "\n*****************"

	@echo "* Download part 2 directly"
	@echo ""
	@echo "make download day=01"

newday:
	echo $(DIRECTORY)
	if [ -d "./$(DIRECTORY)" ]; then \
	    echo "Dir exists"; \
	    exit; \
	fi
	cargo new --bin day$(day)
	cp ./template/src/main.rs ./day$(day)/src/main.rs
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year);

submit:
	aoc s $(part) $(answer) -d $(day) -y $(year) && \
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o

download:
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o
