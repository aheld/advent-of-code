# vim: set noexpandtab:

year=$(shell basename `dirname $(CURDIR)`)
DIRECTORY=$(notdir $(CURDIR))
day=$(subst day,,$(DIRECTORY))

help:
	@echo $(day)

	@echo "* Submit the answer"
	@echo ""
	@echo "make submit part=1 answer=1234"
	@echo "make submit part=2 answer=1234"
	@echo "\n\n"
	@echo "* Download part 2 (After completion of part 1)"
	@echo ""
	@echo "make download"
	@echo "\n\n"
	@echo "* Show the Calendar"
	@echo ""
	@echo "make calendar"
	@echo "\n\n"
	@echo "* Run tests"
	@echo ""
	@echo "make test"
	@echo "make watchtest"

test t:
	cargo test

watchtest wt:
	cargo watch -x "test -q"

submit s:
	aoc s $(part) $(answer) -d $(day) -y $(year) && \
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o

download d:
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o

calendar c:
	aoc c -y $(year)
