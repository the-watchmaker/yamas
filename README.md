# RunYAML
YAML as script language.

## Why?

I use lots of YAML in different apps. They all have their own implied logic behind data structure and documentation. Nowadays, YAML is being used as near-script language anyways. (Also it looks like Python.)


## Syntax Ideation

### "Hello World"

```yaml
THREAD@:
  - ECHO@: "Hello World \n"
  - LINE@: "Hello World"
```

### "if"

```yaml
RUN@:
  - SET_STR@:
      - NAME@: hello
        VALUE@: world
  - IF@:
      AND@:
        - ${hello} == "world"
      THEN@:
        - ECHO@: |
           Hello World! 
```

### "go to"

```yaml
RUN@:
  - SET_STR@:
      - NAME@: hello
        VALUE@: world
  - IF@:
      AND@:
        - ${hello} == "world"
      THEN@:
        - GOTO@: 1
        - GOTO@: 2
        - GOTO@: 3
  LOC_1@:
    - ECHO@: "Hello"
    - RETURN@
  LOC_2@:
    - ECHO@: " World"
    - RETURN@
  LOC_3@:
    - ECHO@: "\n"
    - RETURN@
```
