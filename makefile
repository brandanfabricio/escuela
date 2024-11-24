EXEC = shopy.exe
BUILD_DIR = build
NPM_DIR = ui/shopy
BUILD_CMD = go build -o $(BUILD_DIR)/$(EXEC)
BUILD_CMD = go build -o $(BUILD_DIR)/$(EXEC)
NPM_BUILD_CMD = cd $(NPM_DIR) && bun run build
COPY_CMD = xcopy shopy.db $(BUILD_DIR)
# RM_CMD = $(if $(filter Windows_NT,$(OS)),rmdir /S /Q public,rm -rf public)

all: 
	@echo "Compilando el proyecto..."
	$(BUILD_CMD)
	$(NPM_BUILD_CMD)
	$(COPY_CMD)
	