.PHONY: all
all: cpu_6.prof cpu_7.prof cpu_8.prof cpu_9.prof

cpu_6.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_6.prof -bench=. -measurements=measurements_6.txt

cpu_7.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_7.prof -bench=. -measurements=measurements_7.txt

cpu_8.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_8.prof -bench=. -measurements=measurements_8.txt

cpu_9.prof: io.go main.go main_test.go types.go
	go test -cpuprofile=cpu_9.prof -bench=. -measurements=measurements_9.txt

cpu_6_prof.txt: cpu_6.prof
	go tool pprof -top cpu_6.prof > cpu_6_prof.txt

cpu_7_prof.txt: cpu_7.prof
	go tool pprof -top cpu_7.prof > cpu_7_prof.txt

cpu_8_prof.txt: cpu_8.prof
	go tool pprof -top cpu_8.prof > cpu_8_prof.txt

cpu_9_prof.txt: cpu_9.prof
	go tool pprof -top cpu_9.prof > cpu_9_prof.txt
