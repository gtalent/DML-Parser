include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.pkg

link:
	6l -o parser _go_.6
