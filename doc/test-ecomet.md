# Running Ecomet Query Language (EQL) Tests

This document describes how to run the tests for the Ecomet Query Language (EQL) dialect within the SQLS project.

## Prerequisites

*   **Go Programming Language:** Ensure you have Go installed and configured on your system. You can download Go from the official website: [https://go.dev/dl/](https://go.dev/dl/)
*   **SQLS Project Repository:** You should have the SQLS project repository cloned to your local machine and be in the project's root directory.

## Running the Tests

The EQL dialect tests are located in the `parser` directory of the SQLS project. To run these tests, follow these steps:

1.  **Navigate to the `parser` directory:**
    Open your terminal and use the `cd` command to navigate to the `parser` directory within your SQLS project:

    ```bash
    cd parser
    ```

2.  **Execute the `go test` command:**
    From the `parser` directory, run the following command:

    ```bash
    go test
    ```

    This command will:
    *   Compile the test files in the `parser` directory (including `eql_test.go`).
    *   Run all the test functions defined in these files.
    *   Display the test results in your terminal.

3.  **Interpret the Test Results:**

    *   **Successful Tests:** If all tests pass, you will see output similar to this:

        ```
        PASS
        ok      github.com/sqls-server/sqls/parser   0.023s
        ```
        The exact output might vary slightly depending on your Go module path and system.

    *   **Failed Tests:** If any tests fail, the output will provide details about the failures, including:
        *   The name of the failing test.
        *   The file and line number where the failure occurred.
        *   A description of the expected and actual results that caused the failure.

4.  **Verbose Test Output (Optional):**
    For more detailed output during test execution, use the `-v` flag with the `go test` command:

    ```bash
    go test -v
    ```

    This will print the name of each test function as it is executed, which can be helpful for debugging and understanding the test flow.

## Example Test Run

Here's an example of what a successful test run might look like in your terminal:

```bash
user@computer:~/sqls/parser$ go test
PASS
ok      github.com/sqls-server/sqls/parser   0.025s
user@computer:~/sqls/parser$
```

By following these instructions, you can easily run the EQL dialect tests and ensure that your EQL implementation is working as expected. As you add more features and tests for the EQL dialect, remember to use the `go test` command to verify your changes.
