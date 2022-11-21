# GOBOLT
*A bolted joint strength calculator written in Go*

## Summary
This program calculates strength margins for bolted joints, using methodology compiled from multiple industry references.

More importantly, its development is serving as as my introduction to the go language.

## Current Status
*Alpha 0.0.1*

- 2022-11-20
    - Utilities menu implemented
        - `area` calculates the circular area for a given diameter
        -  `preload` calculates min/nominal/max preload for a bolted joint

See [ROADMAP](ROADMAP.md) for list of planned features.

## References
- [NASA-TM-106943](https://ntrs.nasa.gov/citations/19960012183) "Preloaded joint analysis methodology for space flight systems" (1995)
- [NASA-STD-5020](https://standards.nasa.gov/standard/nasa/nasa-std-5020) "Requirements for Threaded Fastening Systems in Spaceflight Hardware"
- Budynas & Nisbett, "Shigley's Mechanical Engineering Design (8th edition)", 2006, McGraw-Hill, ISBN 978-0073312606 