DEMODIR = ./demo
ALGORITHMDIR = ./algorithms
BUILDOUTPUTDIR = ./src/bin

build_demo: $(DEMODIR)/*
	for file in $^ ; do \
		if [ -f $$file ] && [ "$${file##*.}" = "go" ] ; then \
			file_name=$${file##*/} ; \
			file_name=$${file_name%.*} ; \
			go build -buildmode=c-shared -o $(DEMODIR)/build/$$file_name.so $$file ; \
		fi ; \
	done

build: $(ALGORITHMDIR)/*
	for file in $^ ; do \
		if [ -f $$file ] && [ "$${file##*.}" = "go" ] ; then \
			file_name=$${file##*/} ; \
			file_name=$${file_name%.*} ; \
			go build -buildmode=c-shared -o $(BUILDOUTPUTDIR)/$$file_name.so $$file ; \
		fi ; \
	done

clear:
	rm -rf $(BUILDOUTPUTDIR)/*

clearall:
	rm -rf $(BUILDOUTPUTDIR)/*
	rm -rf $(DEMODIR)/build/*

dist: build
	echo "Creation of the distribution package is not configured yet!"
