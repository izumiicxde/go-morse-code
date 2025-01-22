# Morse Code Decoder & Generator

A Morse code generator and decoder written in Go.

## Usage

```
morse-code-decoder [command]
```

### Available Commands:

- **completion**  
  Generate the autocompletion script for the specified shell.

- **decode**  
  Decode Morse code to text.

- **generate**  
  Generate Morse code from a given text.

- **help**  
  Show help about any command.

### Flags:

- `-h, --help`  
  Display help information for the `morse-code-decoder` command.

- `-t, --toggle`  
  Toggle an optional feature or setting.

## Examples:

1. **Generate Morse Code**  
   To generate Morse code from a text:

   ```
   morse-code-decoder generate -t "HELLO"
   ```

2. **Decode Morse Code**  
   To decode Morse code to text:

   ```
   morse-code-decoder decode -t ".- .... . .-.. .-.. ---"
   ```

## Installation

To install the project, clone this repository and build it using Go:

```
git clone https://github.com/yourusername/morse-code-decoder.git
cd morse-code-decoder
go build
```

Run the built executable:

```
./morse-code-decoder
```
