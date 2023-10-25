# Contributing
in order to contribute to this project, you need to follow the following steps:
1. Create a go module for your folder

    ```bash
    go mod init github.com/kitrak-rev/pykitzoid/<folder_name>
    ```

2. run [golangci-lint](https://github.com/golangci/golangci-lint) to check for errors

    ```bash
    golangci-lint run
    ```

3. once the code is complete and linted, see the steps in the [README.md](../README.md) file to create a pull request