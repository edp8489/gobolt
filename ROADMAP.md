# goBolt CLI Development Roadmap

1. Skeleton CLI
  - Options for "Interactive" and "batch" modes
  - Launching with no commands defaults to a welcome screen asking to pick batch/interactive
  - Setting to track units (inch/metric)
  - All selections print a basic debug message to console
2. CLI development
  1. Batch mode Palette
    - Takes path to an input file
    - Unmarshals inputs into respective structs
    - Runs calculations
    - Prints results to console
    - Option to save to output file
  2. Interactive mode Palette
    - Walks user through defining each input required for a joint
    - Command: Unit selection (imperial or metric)
    - Palette: Joint definition
      - Command: Configuration
      - Number of members
        - Create joint Type with proper array size
      - Member property definition loop
3. Model development
  - Type definition
    - Bolt
    - Joint
    - Member
  - Helper functions
    - Bolt geometry
      - Thread geom
      - Areas
      - Stiffness
    - Joint
      - Member stiffness
    - Preload calcs
    - Margin calcs
      - Bolt tension
      - Bolt shear
      - Bolt shear-tension interaction
      - Joint separation
      - Joint member bearing
    