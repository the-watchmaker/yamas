# RunYAML
YAML as script language.

## Synopsis

Why? I use lots of YAML in different apps. They all have their own implied logic behind data structure and documentation. Nowadays, YAML is being used as near-script language anyways. (Also it looks like Python.)


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
