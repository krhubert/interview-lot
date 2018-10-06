# parking_lot

## Overview

LoT (Language of Tomorrow) is new producto that help solve problem with parking slots (in galleries, merkets and many more)

The whole appliaction is shipped to you with builtin database and great interactive shell!

## Build 

Install `golang` version >=1.11 ([see docs](https://golang.org/doc/install))

```
$ go get parking/lot
```

And it's done.

## Example

Take a look at `example.lot` file

## Language specification

```
registration_number = ^[A-Z]{2}-\d{2}-[A-Z]{1,2}-\d{3,4}$ (example: KA-01-HH-1234)

create_parking_lot INT
park STRING(registration_number) STRING(color)
leave INT
registration_numbers_for_cars_with_colour STRING(color)
slot_numbers_for_cars_with_colour STRING(registration_number)
slot_number_for_registration_number STRING(registration_number)
status
```

## Shell

Start shell with `$ parking_lot` (type `exit` to quit the shell).

## Roadmap

Check out ROADMAP.md in this repository.

## How to contribute?

Check out CONTRIBUTING.md in this repository.
