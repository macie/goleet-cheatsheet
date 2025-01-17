# This Makefile intended to be POSIX-compliant (2018 edition with .PHONY target).
#
# .PHONY targets are used by task definintions.
#
# More info:
#  - docs: <https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html>
#  - .PHONY: <https://www.austingroupbugs.net/view.php?id=523>
#
.POSIX:
.SUFFIXES:


#
# PUBLIC MACROS
#

TEX      = xelatex
GO       = go
BOOKNAME = booklet
TEMPDIR  = /tmp/goleet
SRCDIR   = ./code
DESTDIR  = ./dist
SHA256   = sha256sum -c  # for OpenBSD use: sha256 -c

#
# INTERNAL MACROS
#


#
# DEVELOPMENT TASKS
#

.PHONY: all
all: dist

.PHONY: clean
clean:
	@echo '# Remove temporary files' >&2
	rm -rf $(TEMPDIR)
	@echo '# Delete distribution files' >&2
	rm -rf $(DESTDIR)

.PHONY: info
info:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@echo; $(SHA256) --version || true
	@$(TEX) -version || true

.PHONY: check
check:
	@echo '# Static analysis' >&2
	cd $(SRCDIR) ; $(GO) vet ./...

.PHONY: test
test:
	@echo '# Unit tests' >&2
	cd $(SRCDIR) ; $(GO) test ./...

.PHONY: booklet
booklet: check test $(TEMPDIR)/$(BOOKNAME).pdf
	@echo '# Create distribution directory' >&2
	mkdir -p $(DESTDIR)
	@echo '# Copy booklet' >&2
	cp $(TEMPDIR)/$(BOOKNAME).pdf $(DESTDIR)
	@echo '# Add executable checksum to: $(DESTDIR)/$(BOOKNAME).sha256sum' >&2
	cd $(DESTDIR); sha256sum $(BOOKNAME).pdf >$(BOOKNAME).sha256sum

.PHONY: preview
preview: $(TEMPDIR)/$(BOOKNAME).pdf
	@printf '# Open preview of: %s\n' "$(TEMPDIR)/$(BOOKNAME).pdf" >&2
	xdg-open $(TEMPDIR)/$(BOOKNAME).pdf


#
# DEPENDENCIES
#

$(TEMPDIR)/$(BOOKNAME).pdf: $(BOOKNAME).tex $(SRCDIR)/*.go
	@echo '# Create temporary directory' >&2
	mkdir -p $(TEMPDIR)
	@echo '# Compile LaTeX document' >&2
	@if [ ! -f $(TEMPDIR)/$(BOOKNAME).toc ]; then \
		$(TEX) -interaction=nonstopmode -file-line-error -output-directory="$(TEMPDIR)" $<; \
		echo; echo '# Create table of contents' >&2; \
	fi
	$(TEX) -interaction=nonstopmode -file-line-error -output-directory="$(TEMPDIR)" $<
	@echo
