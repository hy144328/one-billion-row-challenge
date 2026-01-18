SRC_FILES := io.go main.go main_test.go types.go

.PHONY: all
all:
	for i in $$(seq 6 9); do \
		make cpu_run_$${i}; \
	done

.PRECIOUS: cpu_run_%.prof cpu_run_%.log
cpu_run_%.prof cpu_run_%.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run_$*.prof -bench=^BenchmarkRun$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run_$*.log

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
