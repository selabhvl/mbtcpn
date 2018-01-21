# A Go implementation of TPC protocol for the MBT/CPN tool

This Go implementation consists of A test adapter and the coordinator of the TPC protocol.
The test adapter includes a reader and a tester.
The reader is used to read test cases in XML files generated from the MBT/CPN tool.
The tester can use the test cases to execute coordinator (system under test) and compare the test results with oracles.
To run the tests, we assume you have Go development environment on your machine. 

## Run Test
1. Create a new folder named as xml under twophasecommit directory and put the generated xml file (tests.xml) into it,

2. From terminal, and under the tester directory of the adaptor, run tests by the commend: go test -v coordinator_test.go  