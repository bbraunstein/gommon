# logger
Wrapper for go.uber.org/zap with Elastic Common Schema (ECS) default format.

## Installation

```bash
go get -u github.com/bbraunstein/logger
```

## Usage

```golang
import github.com/bbraunstein/logger
```

## Quick Start

In contexts where performance is nice, but not critical, use the `SugaredLogger`. It is 4-10x faster than other structured logging packages and includes both structured and `printf`-style APIs.

```golang
log := logger.NewWithSugaredLogger()
defer log.Sync() // flushes buffer, in any
log.Infow("Failed to fetch URL",
    // Structured context as loosely typed key-value pairs.
    "url", url,
    "attempt", 3,
    "backoff", time.Second,
)
log.Infof("Failed to fetch URL: %s", url)
```

When performance and type safety are critical, use the `Logger`. Its even faster than the `SugaredLogger` and allocates far less, but only supports structured logging.

```golang
log := logger.New()
defer log.Sync()
log.Info("failed to fetch URL",
   // Structured context as strongly typed Field values.
   zap.String("url", url),
    zap.Int("attempt", 3),
    zap.Duration("backoff", time.Second),
)
```