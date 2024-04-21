
define create_file_list
	find /space/$(1) -type f | xargs -n1 basename | sort -n | tee "tmp/$(1).txt"
	wc -l "tmp/$(1).txt"
endef

TRUNK = /space/$(PROJECT)/trunk

scan:
	DEBUG=1 ONLYSCAN=1 $(EXECUTABLE)

ci: build
	mkdir -pv -- "$(TRUNK)"
	DEBUG=1 $(EXECUTABLE)
	$(call create_file_list,test)
	$(call create_file_list,$(PROJECT))
	diff -Nru tmp/test.txt tmp/$(PROJECT).txt
