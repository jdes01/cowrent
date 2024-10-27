BASE_URL := http://localhost:8080

upload:
	@if [ -z "$(UUID)" ]; then \
		read -p "Enter UUID: " UUID; \
	else \
		UUID="$(UUID)"; \
	fi; \
	if [ -z "$(IMAGE_PATH)" ]; then \
		read -p "Enter IMAGE_PATH: " IMAGE_PATH; \
	else \
		IMAGE_PATH="$(IMAGE_PATH)"; \
	fi; \
	curl -X POST "$(BASE_URL)/api/coworking/$$UUID/image" -F "file=@$$IMAGE_PATH"

create-coworking:
	python3 ./scripts/create-coworking.py