<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Beagle Generate Scaffold

## Premisses

Be in a linux environment (Linux or Mac)

## Command

- Prompt

```bash
rit beagle generate scaffold
```

- Docker

```bash
rit beagle generate scaffold --docker
```

## description

This formula generates a backend project with beagle and also, if needed, generates a web and/or mobile project with beagle.

It allows the user to inform 10 different kinds of inputs:

- Backend project name (Text)
- Package name (ex: com.example) (Text)
- JDK version(8+) (default: 13) (Number)
- Kotlin version (1.3+) (Text)
- Beagle version (Text)
- Choose between spring or micronaut (Text)
- Tell if you want to create a Web project (Boolean)
  - Web project name (Text)
- Tell if you want to create Mobile project (Boolean)
  - Mobile project name (Text)

## Demo

![gif](https://github.com/ZupIT/ritchie-formulas/raw/master/beagle/generate/scaffold/doc/beagle-generate-scaffold.gif)
