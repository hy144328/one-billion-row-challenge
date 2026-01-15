SRC_FILES := hash_map.go io.go main.go main_test.go types.go

.PHONY: test
test:
	go test .

.PHONY: full_all
full_all:
	for j in $$(seq 0 6); do \
		make cpu_run$${j}_9; \
	done

.PHONY: cpu_all
cpu_all:
	for i in $$(seq 6 8); do \
		for j in $$(seq 0 6); do \
			make cpu_run$${j}_$${i}; \
		done; \
	done

.PRECIOUS: cpu_run%_6.prof cpu_run%_6.log
cpu_run%_6.prof cpu_run%_6.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run$*_6.prof -bench=^BenchmarkRun$*$$ -run=^$$ -count=10 -in=measurements_6.txt | sed "s/Run$*/Run/g" > cpu_run$*_6.log

.PRECIOUS: cpu_run%_7.prof cpu_run%_7.log
cpu_run%_7.prof cpu_run%_7.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run$*_7.prof -bench=^BenchmarkRun$*$$ -run=^$$ -count=10 -in=measurements_7.txt | sed "s/Run$*/Run/g" > cpu_run$*_7.log

.PRECIOUS: cpu_run%_8.prof cpu_run%_8.log
cpu_run%_8.prof cpu_run%_8.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run$*_8.prof -bench=^BenchmarkRun$*$$ -run=^$$ -count=10 -in=measurements_8.txt | sed "s/Run$*/Run/g" > cpu_run$*_8.log

.PRECIOUS: cpu_run%_9.prof cpu_run%_9.log
cpu_run%_9.prof cpu_run%_9.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run$*_9.prof -bench=^BenchmarkRun$*$$ -run=^$$ -count=1 -in=measurements_9.txt | sed "s/Run$*/Run/g" > cpu_run$*_9.log

.PRECIOUS: cpu_%.svg
cpu_%.svg: cpu_%.prof
	go tool pprof -svg -output=$@ $<

.PRECIOUS: cpu_%.txt
cpu_%.txt: cpu_%.prof
	go tool pprof -top -cum -output=$@ $<

.PRECIOUS: cpu_%.stat
cpu_%.stat: cpu_%.log
	go tool benchstat $< > $@

cpu_%: cpu_%.txt cpu_%.svg cpu_%.stat
	@:

.PHONY: clean
clean:
	-rm cpu_run*.log
	-rm cpu_run*.prof
	-rm cpu_run*.stat
	-rm cpu_run*.svg
	-rm cpu_run*.txt
