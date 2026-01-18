data.csv: parse.py
	python3 parse.py

main.svg: main.py data.csv
	python3 main.py

.PHONY: clean
clean:
	-rm main.svg

.PHONY: distclean
distclean:
	make clean
	-rm data.csv
