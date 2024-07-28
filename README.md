This work is from cloning the [go routines pipeline respoitory](https://github.com/code-heim/go_21_goroutines_pipeline) at code_heim. Additional work was done to add unit testing, benchmark methods to test the if the process runs faster with or without goroutines with custom images.

Clone the repository and type 'go run main.go' in the terminal to run the program.

Changes made:
1. Error checking already existed for file input and output.
2. Added option to run without goroutines depending on user input. Input 1 for running with concurrency, 2 for without concurrency.
3. Added time library to measure throughput time of running with and without concurrency.
Running with concurrency takes around 136 ms, while running without concurrency takes ~182 ms. These times can differ if run on different machines with different images.
4. Added two unit tests.
