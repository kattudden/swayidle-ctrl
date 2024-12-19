# swayidle-ctrl

This Go application allows you to manage a swayidle process for screen and display timeout on Sway desktops.

## Features

- Starts a swayidle process in the background with configurable timeouts.
- Terminates any existing swayidle process before starting a new one.
- Provides a flag to disable screen and display timeout functionalities.

## Usage

This application can be run from the command line. Here's how:

```
swayidle-ctrl -t1 <timeout1> -t2 <timeout2> [-off]

Arguments:
  -t1 (default: 300): Sets the timeout in seconds for screen lock activation.
  -t2 (default: 600): Sets the timeout in seconds for display power off.
  -off (default: false): Disables screen lock and display timeout functionalities.
```

### Example:

Start swayidle with 5 minutes screen lock timeout and 10 minutes display power off timeout:

``` bash
swayidle-ctrl -t1 300 -t2 600
```

Disable screen lock and display timeout:

```bash
swayidle-ctrl -off
```

## Installation

1. Download the source code or clone the repository (if applicable).
2. Run `go build` to create the executable file (`swayidle-ctrl`).
3. Place the executable in a directory accessible from your PATH environment variable.

**Alternatively, you can run the application directly from the source code using `go run main.go`.**

### Additional Notes

* This application uses `sh -c` to execute the swayidle command. You might need to adjust this based on your specific setup.
* The application terminates any existing `swayidle` process before starting a new one. This ensures you don't have multiple instances running concurrently.
* Consider adding support for additional swayidle functionalities or configuration options in future versions.


This readme provides a basic overview of the application. You can customize it further by adding information about:

* Contributing guidelines
* License information
* Troubleshooting tips