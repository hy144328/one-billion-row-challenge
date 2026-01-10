.PHONY: all
all: cpu_6.prof cpu_7.prof cpu_8.prof cpu_9.prof

cpu_%.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=$@ -bench=. -measurements=measurements_$*.txt

cpu_%.txt: cpu_%.prof
	go tool pprof -top -output=$@ $<

cpu_%.svg: cpu_%.prof
	go tool pprof -svg -output=$@ $<
