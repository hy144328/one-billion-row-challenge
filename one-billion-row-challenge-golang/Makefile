.PRECIOUS: cpu_run_%.prof
cpu_run_%.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=$@ -bench=^BenchmarkRun$$ -in=measurements_$*.txt

.PRECIOUS: cpu_%.svg
cpu_%.svg: cpu_%.prof
	go tool pprof -svg -output=$@ $<

.PRECIOUS: cpu_%.txt
cpu_%.txt: cpu_%.prof
	go tool pprof -top -output=$@ $<

cpu_%: cpu_%.txt cpu_%.svg
	@:
