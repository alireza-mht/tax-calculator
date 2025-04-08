# Tax Calculator

This project is designed to calculate income tax based on a provided salary amount for a specific year. It leverages the `ptsdocker16/interview-test-server` Docker image to retrieve tax bracket information for the given year.

## How to Build

A `Makefile` is included to streamline the build process. Below are the primary targets and their purposes:

- **`make deps`**: Installs the project dependencies. The required dependencies are:
  1. **oapi-codegen**: Generates OpenAPI requirements for the project.
  2. **golangci-lint**: Analyzes the code for linting, identifying potential errors and improving code quality.
- **`make build`**: Compiles the application, generating a binary in the `bin/` folder at the project root.
- **`make lint`**: Displays the linter output to review code quality and potential issues.
- **`make clean`**: Removes generated files from the build process.

## How to Run

Before running the application, you must pull and start the `ptsdocker16/interview-test-server` locally on port `5001`. Use the following commands:

```bash
docker pull ptsdocker16/interview-test-server
docker run --init -p 5001:5001 -it ptsdocker16/interview-test-server
```

After that it's possible to run the command below to run the server from root project:
```bash
./bin/tc serve
```
The default value for host and port is `localhost:8383`, but it can be customized if needed.

The server will run, and you can make requests to it. An example of a request is provided below:

```bash
$ curl -X GET http://localhost:8383/v1/tax-calculator/tax-year/2023?salary=150000 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   222  100   222    0     0     54      0  0:00:04  0:00:04 --:--:--    54
{
  "effectiveRate": 0.20130546,
  "taxPerBand": [
    {
      "max": 53359,
      "min": 0,
      "rate": 0.15,
      "tax": 8003.85
    },
    {
      "max": 106717,
      "min": 53359,
      "rate": 0.205,
      "tax": 10938.39
    },
    {
      "max": 150000,
      "min": 106717,
      "rate": 0.26,
      "tax": 11253.579
    }
  ],
  "totalTax": 30195.82
}
```
# Features

1. **Automated CI Build**: For this project, a GitHub Actions workflow is utilized, incorporating several stages.
2. **Error Handling**: Errors are managed uniquely using error codes provided in the common package.
3. **Unit Tests**: Unit tests are available for certain functions that do not require mocking.
4. **Logging**: A logging package is included to ensure proper logging of essential information.
5. **API Design**: The general API design is available under `./api/openapi.yml`.
6. **Retry Mechanism**: The current retry mechanism employs a ticker to attempt retries at fixed intervals. However, this interval is not consistently maintained because it fails to account for processing time. A more precise approach would be to utilize the time module and implement proper sleep durations to ensure accurate retry timing.


# Wish List

Due to limited time for completing the features, the following are key items on the wish list for future improvements:

1. **Add Dockerfile for Application**: To create an image for the application, a Dockerfile would be required.
2. **Provide a Helm Chart**: A Helm chart would be beneficial for deploying the application in a Kubernetes environment.
3. **Integration Testing**: It’s feasible to run the current application and an external service within the CI pipeline, testing the provided endpoint by running both containers in the pipeline.
4. **Component Testing**: Since the application depends on an external service, component testing would ensure that the components function together effectively, such as by mocking the external service.
5. **Documentation**: Cobra provides CLI documentation that offers a helpful overview of existing commands. Additionally, Go’s built-in documentation is valuable. Including these in the project would enhance interaction with the application.
6. **Complete Makefile**: There are some steps remaining in the Makefile that would be nice to have. Completing targets like `make test` and `make doc` would improve usability.