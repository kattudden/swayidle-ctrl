# swayidle-ctrl

This Go application allows you to enable or disable the swayidle service.
It would also be possible to do it directly with systemctl...

## Features

- stop / start the swayidle service.

## Usage

This application can be run from the command line. Here's how:

```txt
swayidle-ctrl [-on|-off]

Arguments:
  -on (default: false): Enables screen lock and display timeout functionalities.
  -off (default: false): Disables screen lock and display timeout functionalities.
```

### Examples

Enable screen lock and display timeout:

``` bash
swayidle-ctrl -on
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

- This application uses `sh -c` to execute the swayidle command. You might need to adjust this based on your specific setup.
- The application terminates any existing `swayidle` process before starting a new one. This ensures you don't have multiple instances running concurrently.
- Consider adding support for additional swayidle functionalities or configuration options in future versions.

This readme provides a basic overview of the application. You can customize it further by adding information about:

- Contributing guidelines
- License information
- Troubleshooting tips
