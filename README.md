# RunYAML
YAML as script language.

## Ideation and why?
I use lots of YAML in different apps. They all have their own implied logic behind data structure and documentation. Nowadays, YAML is being used as near-script language anyways. (Also it looks like Python.)


## Syntax Ideation

### "Hello World"

```yaml
thread@:
  - echo@: "Hello World \n"
  - lineEcho@: "hello world"
```

### "if"

```yaml
run@:
  - set@:
      - type@: string
        name@: hello
        value@: world
  - if@:
      con@:
        - left@: ${hello}
          con@: ==
          right@: world
      then@:
        - thread@:
            - echo@: |
                Hello World! 
```
