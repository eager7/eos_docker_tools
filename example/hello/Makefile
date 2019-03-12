TARGET := hello.wasm
SRC := hello.cpp

CC := eosio-cpp
RM := -rm

.PHONY: all clean distclean test
all: $(TARGET)

$(TARGET):$(SRC)
	$(CC) $^ -o $@

clean:
	$(RM) *.wasm

test: $(SRC)
	@echo $(CC)
	@echo $(TARGET)
	@echo $^
	@echo $@
