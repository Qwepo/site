.PHONY: run_db

run_db:
	docker run -d -p9088:9088 -p6534:6534 -it reindexer/reindexer 

