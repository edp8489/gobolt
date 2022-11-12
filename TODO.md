2022-11-07
- Start by fleshing out Utils palette
  - Prompt to set precision when printing calculated values
  - ~~Area calculator~~ (done)
  - Pitch, minor diameter calcs for given basic diam, thread pitch, class
  - Thread shear area
  - Preload calcs (in work)
  - Refactor prompt helpers to public package
  - Switch from `promptui` library to the better documented `go-survey`
  - Implement `bcicen/go-units` for unit consistency on inputs


2022-07-02
- Refactor component type definitions into individual files
- Make type definitions private; create public factory methods to return new instances with appropriate default values and derived properties
- Database of allowable diameter-pitch combinations (JSON data file?)
- Create methods for each derived property
  - Operate on pointer to type
  - Minor diameter, pitch diameter
    - Metric: Shigley, table 8-1 footnote
    - Imperial: Shigley Table 8-2 footnote
  - Area
    - FED-STD-H28/2B, Table II.B.1, eq 1(a), 1(b)
    - Shank area based on nominal diameter
    - Tensile area at midpoint between pitch and minor diameter
  - Thread Shear Area
    - (maybe) type switch to change calc based on internal/external thread
    - Internal: FED-STD-H28/2B, Table II.B.1, eq 2
    - External: FED-STD-H28/2B, Table II.B.1, eq 4
- Package for margin calc functions
  - Fastener
    - Shear margin
    - Tension margin
    - Bending margin
    - Shear-Tension interaction
    - Shear-Tension-Bending interaction
    - Thread shear margin
  - Joint members
    - Bearing margin
  - Preload analysis