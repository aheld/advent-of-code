default: help

year?=$(notdir $(CURDIR))

DIRECTORY=day$(day)

help:
	@echo "\n\n*****************"
	@echo "* Download a new puzzle file and create a directory"
	@echo "* with a rust project"
	@echo ""
	@echo "make newday day=01"

	@echo "\n\n*****************"
	@echo "* Show the Calendar"
	@echo ""
	@echo "make calendar"

newday:
	if [ -d "./$(DIRECTORY)" ]; then \
	    echo "Dir exists"; \
	    exit; \
	fi
	cargo new --bin day$(day)
	cp ./template/src/main.rs ./day$(day)/src/main.rs
	cp ./template/Makefile ./day$(day)/Makefile
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year);

calendar c:
	aoc c -y $(year)
