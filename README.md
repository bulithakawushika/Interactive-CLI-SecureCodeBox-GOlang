# Go CLI Tool for Creating SecureCodeBox Resource Files

This repository contains a simple Go language CLI tool that allows users to create SecureCodeBox resource files by providing inputs through the command line interface (CLI). The tool creates resource files based on the SecureCodeBox resource format, specifically designed for executing Nmap scans.

## Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/bulithakawushika/Create-YAML-file-using-CLI-SecureCodeBox.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Create-YAML-file-using-CLI-SecureCodeBox
    ```

3. Build the CLI tool:

    ```bash
    go build -o scb-create-resource
    ```

## Usage

To use the CLI tool, execute the binary `scb-create-resource`:

```bash
./scb-create-resource
```

Follow the prompts to provide necessary inputs for creating the SecureCodeBox resource file. The tool will generate a `.yaml` file in the specified location within the project folder.

### Scan Types

The tool currently supports the following scan types:

1. **Scanning network**: Allows you to perform Nmap scans on specified target URLs.

    ```bash
    Select Scan Type:
    1- Scanning network
    2- Repeat Scans on a Schedule
    3- Scanning Web Application
    4- Post-processing with Hooks
    5- Storing Findings with Persistence Hooks
    6- Automatically Scan your Cluster with Autodiscovery
    7- Enforce Engagement
    Enter your choice: 1
    Enter target URL: example.com
    ```

    This will generate a YAML file named `output_<timestamp>.yaml` in the `output_yaml` directory, containing the Nmap scan configuration.

## YAML File Structure

The generated YAML files follow the structure below:

```yaml
apiVersion: execution.securecodebox.io/v1
kind: Scan
metadata:
  name: <scan-name>
spec:
  scanType: "1"
  parameters:
    - <target-URL>
```

- `apiVersion`: Indicates the API version used by SecureCodeBox.
- `kind`: Specifies the type of resource, which in this case is a "Scan".
- `metadata.name`: Represents the name of the scan, typically prefixed with "nmap-" followed by the target URL.
- `spec.scanType`: Represents the type of scan being performed.
- `spec.parameters`: Contains parameters specific to the scan type, such as the target URL.

## SecureCodeBox Resource Format

The generated YAML files adhere to the SecureCodeBox resource format. For more information on the SecureCodeBox resource format and available parameters, refer to the [SecureCodeBox documentation](https://www.securecodebox.io/docs/scanners/nmap).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributors

- [Bulitha kawushika](https://github.com/bulithakawushika)
