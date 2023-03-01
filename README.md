<img src="https://user-images.githubusercontent.com/4682613/222035515-8ff9e540-aa55-48a2-8208-48f7699233aa.jpg" width="260" />

Fujiyama[^1]


# Yamas
YAMAS stand for Yet Another Markup As Script


## Why?

I use lots of YAML in different apps. They all have their own implied logic behind data structure and documentation. Nowadays, YAML is being used as near-script language anyways. (Also it looks like Python.)

## Design Principle
1. Super simple
2. No native loops
3. Limited native functions:
   - ECHO@
   - LINE@
   - HTTP@
   - BASH@

## Syntax

### "Hello World"

```yaml
THREAD@:
  - ECHO@: "Hello World \n"
  - LINE@: "Hello World"
```

### "IF"

```yaml
RUN@:
  - SET@:
      - NAME@: hello
        VALUE@: "world"
  - SET@
      - NAME@: peace
      - VALUE@: true
  - IF@:
      AND@:
        - EQ@: 
          - ${hello}
          - "world"
        - EQ@:
          - ${peace}
          - true
      THEN@:
        - ECHO@: |
           Hello World! 
```

### "GO TO"

```yaml
RUN@:
  - SET@:
      - NAME@: hello
        VALUE@: world
  - IF@:
      - EQ@: 
        - ${hello}
        - "world"
      THEN@:
        - GOTO@: 1
        - GOTO@: 2
        - GOTO@: 3
  - LOC_1@:
      - ECHO@: Hello
      - RETURN@
  - LOC_2@:
      - ECHO@: " World"
      - RETURN@
  - LOC_3@:
      - ECHO@: "\n"
      - RETURN@

```


### FUNCTION

```yaml
RUN@:
  - FN@: World
    TAKE@:
      - greeting
    RUN@:
      - IF@: ${greeting} == "world"
        THEN@:
          - RETURN: "World"
  - ECHO@:
      FROM@:
        - "Hello"
        - FROM_CALL@: World
          PUT@:
            - greeting: "Hello"
```

## Statements & declarations

- FROM@
```yaml
RUN@:
  - ECHO@:
      FROM@:
        - "Hello"
        - " World"
        - "\n"

```



## Expressions & operators
- EQ@, NE@, GT@, L@, GTE@, LTE@
```yaml
EQ@:
   - ${foo}
   - "bar"
```

```Javascript
 foo === "bar"
```

- ADD@, SUB@, MUL@, DIV@
```yaml
@SET_NUM:
   - "added"
   - @FROM:
      - @ADD
         - ${added}
         - 1
```

```Javascript
 added = added + 1;
```





## Some random reading
- [How to Build a New Programming Language](https://pgrandinetti.github.io/compilers/page/how-to-build-a-new-programming-language/) by 
- [Designing a Programing Language](http://ducklang.org/designing-a-programming-language-i) By DuckLang.org
- [Programming Language Design](https://bootcamp.uxdesign.cc/programming-language-design-a649513dbcf7) by Saeed Mohajeryami
- [Syntax Design](https://cs.lmu.edu/~ray/notes/syntaxdesign/) by LMU CS
- [Creating A Programming Language From Scratch](https://medium.com/swlh/creating-a-programming-language-from-scratch-244b88e33e2f) by Joao Zsigmond

[^1]: Photo by <a href="https://unsplash.com/@jmeguilos?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Jules Marvin Eguilos</a> on <a href="https://unsplash.com/photos/O3oQg9CPy1k?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>

