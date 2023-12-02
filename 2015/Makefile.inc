submit s:
	aoc s $(part) $(answer) -d $(day) -y $(year) && \
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o

download d:
	cd $(DIRECTORY) && aoc d -d $(day) -y $(year) -o

calendar c:
	aoc c -y $(year)

