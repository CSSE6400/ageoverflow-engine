# Engine

A command line interface for processing the requested content.

**Note**: This project is for use in the cloud infrastructure assessment of the course CSSE6400 Software Architecture at the University of Queensland and is not intended for any other usage.

## Installation

This package must be installed by collecting the static binary given in the releases.

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
    "QXJjaGl0ZWN0dXJlIGlzIGFib3V0IHRoZSBpbXBvcnRhbnQgc3R1ZmYuIFdoYXRldmVyIHRoYXQgaXMu",
    "YSBzb2Z0d2FyZSBhcmNoaXRlY3Qgd2hvIGNvZGVzIGlzIGEgbW9yZSBlZmZlY3RpdmUgYW5kIGhhcHBpZXIgYXJjaGl0ZWN0"
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
    "checksum": "ABCDEF123456",
    "generations": {
      "silent": 0,
      "baby_boomers": 5,
      "x": 15,
      "y": 40,
      "z": 30,
      "alpha": 10
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
