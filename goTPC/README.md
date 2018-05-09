# A Go implementation of TPC protocol for the MBT/CPN tool

This Go implementation consists of a test adapter and the coordinator of the TPC protocol.
The test adapter includes a reader (reader.go) and a tester (coordinator_test.go).
The reader is used to read test cases in XML files generated from the MBT/CPN tool.
The tester can use the test cases to execute coordinator (system under test) and compare the test results with oracles.
To run the tests, we assume you have Go development environment on your machine.

## Run Test
1. under xml folder, there is an example xml file generated to test coordinator

2. To execute the test, from terminal, and under goTPC directory, run the commend: go test -v

3. To run statement coverage, from terminal, run coverage.sh
