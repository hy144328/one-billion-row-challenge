.PHONY: test
test:
	go test .

.PHONY: full_all
full_all:
	for j in "" $$(seq 1 5); do make cpu_run$${j}_9; done

.PHONY: cpu_all
cpu_all:
	for i in $$(seq 6 8); do for j in "" $$(seq 1 5); do make cpu_run$${j}_$${i}; done; done

.PRECIOUS: cpu_run_%.prof cpu_run_%.log
cpu_run_%.prof cpu_run_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run_$*.prof -bench=^BenchmarkRun$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run_$*.log

.PRECIOUS: cpu_run1_%.prof cpu_run1_%.log
cpu_run1_%.prof cpu_run1_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run1_$*.prof -bench=^BenchmarkRun1$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run1_$*.log

.PRECIOUS: cpu_run2_%.prof cpu_run2_%.log
cpu_run2_%.prof cpu_run2_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run2_$*.prof -bench=^BenchmarkRun2$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run2_$*.log

.PRECIOUS: cpu_run3_%.prof cpu_run3_%.log
cpu_run3_%.prof cpu_run3_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run3_$*.prof -bench=^BenchmarkRun3$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run3_$*.log

.PRECIOUS: cpu_run4_%.prof cpu_run4_%.log
cpu_run4_%.prof cpu_run4_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run4_$*.prof -bench=^BenchmarkRun4$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run4_$*.log

.PRECIOUS: cpu_run5_%.prof cpu_run5_%.log
cpu_run5_%.prof cpu_run5_%.log: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_run5_$*.prof -bench=^BenchmarkRun5$$ -run=^$$ -count=10 -in=measurements_$*.txt > cpu_run5_$*.log

.PRECIOUS: cpu_%.svg
cpu_%.svg: cpu_%.prof
	go tool pprof -svg -output=$@ $<

.PRECIOUS: cpu_%.txt
cpu_%.txt: cpu_%.prof
	go tool pprof -top -output=$@ $<

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
