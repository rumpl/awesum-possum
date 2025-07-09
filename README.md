# Awesum Possum

A Golang command-line application that makes streaming calls to the OpenAI API.

## Features

- Simple command-line interface
- Streams responses in real-time
- Uses OpenAI's GPT-3.5 Turbo model by default

## Installation

### Prerequisites

- Go 1.16 or higher
- OpenAI API key

### Build from source

```bash
# Clone the repository
git clone https://github.com/user/awesum-possum.git
cd awesum-possum

# Build the application
go build -o awesum-possum

# Optional: Move to a directory in your PATH
sudo mv awesum-possum /usr/local/bin/
```

### Using Make

```bash
make build    # Build the application
make install  # Install to /usr/local/bin
make clean    # Remove build artifacts
```

## Usage

First, set your OpenAI API key:

```bash
export OPENAI_API_KEY=your_api_key_here
```

Then run the application:

```bash
awesum-possum "What is the capital of France?"
```

## License

MIT License