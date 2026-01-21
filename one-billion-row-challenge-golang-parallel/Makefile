SRC_FILES := io.go main.go main_test.go types.go

.PHONY: test
test:
	go test .

.PRECIOUS: cpu_run_%.prof cpu_run_%.log
cpu_run_%.prof cpu_run_%.log: $(SRC_FILES)
	go test -cpuprofile=cpu_run_$*.prof -bench=^BenchmarkRun$$ -run=^$$ -count=10 -in=measurements_9.txt > cpu_run_$*.log

.PRECIOUS: cpu_run_%.svg
cpu_run_%.svg: cpu_run_%.prof
	go tool pprof -svg -output=$@ $<

.PRECIOUS: cpu_run_%.txt
cpu_run_%.txt: cpu_run_%.prof
	go tool pprof -top -cum -output=$@ $<

.PRECIOUS: cpu_run_%.stat
cpu_run_%.stat: cpu_run_%.log
	go tool benchstat $< > $@

cpu_run_%: cpu_run_%.txt cpu_run_%.svg cpu_run_%.stat
	@:

.PHONY: clean
clean:
	-rm cpu_run_*.log
	-rm cpu_run_*.prof
	-rm cpu_run_*.stat
	-rm cpu_run_*.svg
	-rm cpu_run_*.txt
