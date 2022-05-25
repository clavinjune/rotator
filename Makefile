# Copyright 2022 clavinjune/rotator
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include tools.mk

check:
	@go run $(licenser) verify
	@go run $(linter) run

fmt:
	@gofmt -w -s .
	@go run $(importer) -w .
	@go vet ./...
	@go mod tidy
	@go run $(licenser) apply -r "clavinjune/rotator" 2> /dev/null
test:
	@go test -v -json -coverprofile=coverage.out -covermode=count ./... > result.json.out
test/coverage: test
	@go tool cover -html=coverage.out
