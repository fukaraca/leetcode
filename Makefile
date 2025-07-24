.PHONY: new
new:
ifndef DIFF
	$(error difficulty not set. Usage: make new DIFF=<easy|medium|hard> NAME=<problem-name>)
endif
ifndef NAME
	$(error name not set. Usage: make new DIFF=<easy|medium|hard> NAME=<problem-name>)
endif
	@mkdir -p $(DIFF)/$(NAME)
	@printf "package main\n\nfunc main() {\n}\n" > $(DIFF)/$(NAME)/main.go
	@cd $(DIFF)/$(NAME) && go mod init $(NAME)