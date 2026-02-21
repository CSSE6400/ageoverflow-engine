# Engine

A command line interface for processing the requested content.

> [!NOTE]
> This project is for use in the cloud infrastructure assessment of the course CSSE6400 Software Architecture at the University of Queensland and is not intended for any other usage.


## Installation

This package must be installed by collecting the static binary given in the releases.

### Docker

For installing the latest version inside a dockerfile the following can be used.

```dockerfile
FROM ubuntu:latest

ARG ENGINE_VERSION=v0.9.0

RUN apt-get update && apt-get install -y curl

RUN ARCH=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && \
     curl -o engine -L https://github.com/CSSE6400/ageoverflow-engine/releases/download/${ENGINE_VERSION}/engine-${ENGINE_VERSION}-linux-${ARCH} && chmod +x engine

CMD ["./engine", "--help"]
```

## Usage

#### Help

```bash
engine --help
```

```bash
engine compute --help
```


#### Scanning a request

Reports are generated from a JSON file like so:

```json
{
  "id": "ABCD-1234",
  "content": [
    "MHwyOXwwfDB8MHwxMDB8MHwwfE1pbGxpZSB0aGUgbGFzdCBvZiB0aGUgbWlsbGVubmlhbHMuIA==",
    "MHwyOXwwfDB8MHw2MHw0MHwwfE1pbGxpZSB0aGUgbGFzdCBvZiB0aGUgbWlsbGVubmlhbHMgb3IgdGhlIGZpcnN0IG9mIHRoZSBab29tZXJzPy4g",
    "MHwyOXwwfDB8MTB8MzV8NTV8MHxNaWxsaWUgdGhlIGZpcnN0IG9mIHRoZSBtaWxsZW5uaWFscy4g"
  ]
}
```

The output is a path with a filename but no extension where a .json will be generated.

```bash
engine compute --input examples/input.json --output examples/output
```

You can also use stdin and stdout for input and output.

For Input:

```bash
cat examples/input.json | engine compute --input '-' --output examples/output
```

or 

```bash
cat examples/input.json | engine compute --output examples/output
```

For Output:

```bash
engine compute --input examples/input.json --output '-' > examples/output.json
```

or 

```bash
engine compute --input examples/input.json > examples/output.json
```


Example Output:

```json
{
  "id": "ABCD-1234",
  "results": {
    "checksum": "0x57",
    "generations": {
      "silent": 0,
      "baby_boomers": 0,
      "x": 3,
      "y": 65,
      "z": 31,
      "alpha": 0
    },
    "primary_generation": "y",
    "age": 29
  }
}
```

## Motivation

This project was created for use in the cloud infrastructure assessment of the course CSSE6400 Software Architecture at the University of Queensland.
It is intended to generate an output that requires work, this version accomplishes this by computing an arbitrary BCRYPT hash which is thrown away.
The program then generates a report based on the information given to it.

## Fingerprint

The fingerprint is a given pipe seperated seed for the engine to generate a report.

- fingerprint[0]: iterations of the BCRYPT hash, recommended to be between 8 -> 20 iterations, if the value is 0 then the BCRYPT is skipped.
- fingerprint[1]: value of Age that should be outputted, if multiple requests are sent then these are averaged.
- fingerprint[2]: value of generations.silent, if multiple requests are sent then these are then normalised.
- fingerprint[3]: value of generations.baby_boomers, if multiple requests are sent then these are then normalised.
- fingerprint[4]: value of generations.x, if multiple requests are sent then these are then normalised.
- fingerprint[5]: value of generations.y, if multiple requests are sent then these are then normalised.
- fingerprint[6]: value of generations.z, if multiple requests are sent then these are then normalised.
- fingerprint[7]: value of generations.alpha, if multiple requests are sent then these are then normalised.
- fingerprint[8]: Filler text if needed.

## Checksum

The checksum is the fingerprint[1] added together for all the inputs and then converted to hex.

## Performance Characteristics

These stats were made on a **t3.small** using:

```bash
psrecord "engine ....." --log activity.txt --plot performance.png
```

| Type                      | Stats                       |
|---------------------------|-----------------------------|
| Scan (sm) [12 iterations] | ![](performance/small.png)  |
| Scan (md) [16 iterations] | ![](performance/medium.png) |
| Scan (lg) [18 iterations] | ![](performance/large.png)  |

## Contributing

Contributions are welcome but the project is for the usage in an assessment so some aspects of the program are intentional to create load on the system.

## Changes

### No Releases

No releases have been made.
