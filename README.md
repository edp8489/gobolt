# GOBOLT
*A bolted joint strength calculator written in Go*

## Summary
This program calculates strength margins for bolted joints, using methodology compiled from multiple industry references.

More importantly, its development is serving as my introduction to the *go* language.

## Current Status
*Alpha 0.0.1*

### 2022-11-20
- Utilities menu implemented
- `area` calculates the circular area for a given diameter
-  `preload` calculates min/nominal/max preload for a bolted joint

See [ROADMAP](ROADMAP.md) for a full list of planned features currently in development.

## Usage
Launching the base program by double-clicking the executable or calling `gobolt` without any additional arguments will print the help message and prompt the user to select from a list of available commands.

    $ gobolt
    goBolt is a bolted joint strength calculator.

    This application can be run several ways:

    1. Batch mode to solve predefined input file(s)
    2. Interactive mode that walks the user through
    all aspects of property definition.
    3. Individual calculations via the standalone utils palette

    Usage:
        gobolt [command]

    Available Commands:
        demo        Runs a demo solve
        utils       Palette containing various utility functions

Since only a subset of utilities have been implemented, use `gobolt utils` to launch the program and navigate directly to the list of available utilities.

    $ gobolt utils
    Utility functions and calculators that can be used independently from full joint definition and margin calculations.

    Usage:
        gobolt utils [flags]
        gobolt utils [command]

    Available Commands:
        area        Calculate the circular area for a given diameter
        preload     Calculate preload for a bolted joint

## References
- [NASA-TM-106943](https://ntrs.nasa.gov/citations/19960012183) "Preloaded joint analysis methodology for space flight systems" (1995)
- [NASA-STD-5020](https://standards.nasa.gov/standard/nasa/nasa-std-5020) "Requirements for Threaded Fastening Systems in Spaceflight Hardware"
- Budynas & Nisbett, "Shigley's Mechanical Engineering Design (8th edition)", 2006, McGraw-Hill, ISBN 978-0073312606 