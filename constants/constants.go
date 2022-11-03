package constants

// DefaultLogRotation sets default behavior log rotation
const DefaultLogRotation = false

// MaxFileCount is the limit for maximum files that can be created by the tool
const MaxFileCount = 30

// DefaultFileCount is the default number of files the generator creates.
const DefaultFileCount = 1

// DefaultMetricsServerPortal is the default port on which the prometheus metrics server starts
const DefaultMetricsServerPort = "8080"

//DefaultLineMaxLength is the default max line length
const DefaultLineMaxLength = 1000 //1 KiloBytes

//DefaultLineMinLength is the default min line length
const DefaultLineMinLength = 10 //10 Bytes

//DefaultMaxFileSize is the maximum size in MBytes after which the file will be rotated.
const DefaultMaxFileSize = 100

//DefaultDurationInSeconds the default time the test would run if no duration is specified.
const DefaultDurationInSeconds = 300 //5mins

//DefaultShutdownSeconds the default time the test would run if no duration is specified.
const DefaultShutdownSeconds = 1 //1 second

//DefaultReplayCount is the default replay count if the duration is not specified and the replay count is not specified.
const DefaultReplayCount = 0 //5mins

//DeafaultLinesPerSecond is the default LPS when it is not explicitly specified.
const DefaultLinesPerSecond = 1000 // 1 Lps

//DefaultMultiLinePercentage is the default % of lineCount > 1 when the percentage is not explicitly defined.
const DefaultMultiLinePercentage = 0 //0%

//DefaultMultiLineCount is the default multi-line count when it is not explicitly defined.
const DefaultMultiLineCount = 1 //1 line

//DefaultMaxFileRotationSize is max file size in MB after which the file will be rotated.
const DefaultMaxFileRotationSize = 500

const TestOutputFileRootPath = "/tmp"
