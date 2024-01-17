# A Minimal Logging Library for Go Applications

This library offers a straightforward and adaptable way to log messages with varying severity levels to different destinations, such as files.


## Features

- Log messages with severity levels: Debug, Info, Warning, Error.
- Customize log format: Include timestamp, log level, and message.
- Easy configuration and usage: Designed for seamless integration into Go programs.
- Support for logging to files: Future expansion planned for additional destinations.
- Log filtering (coming soon): Filter based on log level, source module, and content.



## Installation

To incorporate this library into your Go project, run the following command in your terminal:

```go
go get github.com/bitsexplained/logger
```

## Usage
- Here's an example demonstrating how to utilize this library in your Go application:

```go
   package main

   import (
      "github.com/bitsexplained/logger"
   )

    func main() {
        // Create a new logger instance
        logger := logging.NewLogger()

        // Set log level to Debug
        logger.SetLogLevel(logging.Debug)

        // Set log format
        logFormat := logging.LogFormat{
            Timestamp: "2023-04-18 21:54:05",
            Level:     logging.Level,
            Message:   logging.Message,
        }
        logger.SetLogFormat(logFormat)



        // Add file log location
        // ... (similarly, you can add other log location, such as a file or a database)
        logLocation := logging.NewFileLog("logs.txt") // Specify the file name for logs
	    logger.AddLogLocation(logLocation)

        // Log messages
        logger.Log(logging.Debug, "This is a debug message")
        logger.Log(logging.Info, "This is an info message")
        logger.Log(logging.Warning, "This is a warning message")
        logger.Log(logging.Error, "This is an error message")


}

```
## Current Limitations:
- Supports logging to files only.
- Filtering functionality is not yet implemented.

## Future Expansion:
Planned enhancements based on user feedback:
- Additional logging destinations (console, databases, remote logging services)
- Log filtering capabilities



## Contributing
Your contributions are highly valued! To report bugs, suggest features, or contribute code, please create an issue or submit a pull request on the GitHub repository.

## License
his library is freely available under the MIT License.
See  [LICENSE](LICENSE) for details.

