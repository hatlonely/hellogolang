codegen:
	make -C api/echo_proto/
	make -C api/counter_proto/
	cp -r api vendor/
clean:
	make clean -C api/echo_proto/
	make clean -C api/counter_proto/
