include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $(GOFILES)
link: package
	$(O)l -o main _go_.$(O)
