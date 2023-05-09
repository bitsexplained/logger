# A minimal logging library

This is a logging library that provides a simple and flexible way to log messages with different severity levels written to different locations, such as files.


## Features

- Log messages with different severity levels (Debug, Info, Warning, Error)
- Customize log message format with timestamp, log level and message
- Easy to configure and use in any Go program
- Support for logging to files for now - will add support for more logging locations soon
- Filter log messages based on log level, source module, and content - coming soon


## Installation

To use the logging library in your Go project, you can import it using the following command:

```go
import "https://github.com/bitsexplained/logger.git"
```

## Usage
- Here's an example of how you can use logging library in your Go application:

```go
   package main

   import (
      "https://github.com/bitsexplained/logger.git"
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
- In this example, we import the logging library and create a new logger instance. We then configure the logging settings, such as log level and log format. We add a file log location using the NewFileLog method, specifying the file name for logs. Finally, we use the logger to log messages with different severity levels and content using the Log method, and the logged messages will be written to the configured log file for further analysis and debugging.

- Please note that currently, the logging library only supports writing logs to files, but in the future, it may add more logging destinations, such as console, databases, or remote logging services, based on user feedback and requirements.



## Contributing
Contributions to this library is welcomed! If you find a bug, have a feature request, or want to contribute code, please open an issue or submit a pull request to the GitHub repository.

## License
rate limiter library is released under the MIT License.
See  [LICENSE](LICENSE) for details.

