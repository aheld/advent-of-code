# vim: set noexpandtab:

year=$(shell basename `dirname $(CURDIR)`)
DIRECTORY=$(notdir $(CURDIR))
day=$(subst day,,$(DIRECTORY))

help:
	@echo $(day)

	@echo "\n\n*****************"
	@echo "* Download a new puzzle file and create a directory"
	@echo "* with a rust project"
	@echo ""
	@echo "make newday day=01 year=2015"

	@echo "\n*****************"

	@echo "* Download part 2 directly"
	@echo ""
	@echo "make download"
	@echo "* Show the Calendar"
	@echo ""
	@echo "make calendar"

test t:
	cargo test

watchtest wt:
	cargo watch -x "test -q"

include ../Makefile.inc
