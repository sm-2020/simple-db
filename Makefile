build:
	go build simplelb.go

clean:
	rm -rf simplelb
run:
	docker-compose up
present:
	present preso/simple-load-balancer.slide
