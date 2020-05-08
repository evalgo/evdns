.PHONY: all
all: test

test:
	go get -v github.com/jstemmer/go-junit-report
	go build -o go-junit-report github.com/jstemmer/go-junit-report
	go get -v
	go test -v -run=Test_Unit 2>&1 | ./go-junit-report > report.xml

cli:
	GOOS=linux GOARCH=amd64 go build -o evdns-cli.linux.amd64 cmd/evdns-cli/main.go
	GOOS=darwin GOARCH=amd64 go build -o evdns-cli.darwin.amd64 cmd/evdns-cli/main.go
	GOOS=windows GOARCH=amd64 go build -o evdns-cli.windows.amd64.exe cmd/evdns-cli/main.go

service:
	GOOS=linux GOARCH=amd64 go build -o evdns.linux.amd64 cmd/evdns-service/main.go
	GOOS=darwin GOARCH=amd64 go build -o evdns.darwin.amd64 cmd/evdns-service/main.go
	GOOS=windows GOARCH=amd64 go build -o evdns.windows.amd64.exe cmd/evdns-service/main.go

.PHONY: clean 
clean:
	rm -fv evdns-cli.*.amd64 evdns-cli.*.amd64.exe
	rm -fv evdns.*.amd64 evdns.*.amd64.exe
	find . -name "*~" | xargs rm -fv
	rm -fv go-junit-report report.xml

