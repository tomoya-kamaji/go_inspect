
# mockgen
MOCK_SOURCE_DIR=domain
MOCK_DIRDEST_DIR=mocks/repository

mockgen:
	mockgen -source=$(MOCK_SOURCE_DIR)/bookRepository.go -destination=$(MOCK_DIRDEST_DIR)/mockBookRepository.go -package=mocks
