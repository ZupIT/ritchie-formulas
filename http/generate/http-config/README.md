<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD029 -->

# Description

This formula is a helper for ritman http load testing formula, it generates a default configuration template, like the one bellow:

1. The above command will generate a config file named "ritman-target.json" which is pretty much the configuration as showed bellow:

```json
{
  "target": "https://postman-echo.com/post",
  "method": "POST",
  "headers": {
    "Accept": "application/json; charset=utf-8",
    "Accept-Encoding": "gzip, deflate, br",
    "Connection": "Keep-alive",
    "Content-Type": "application/json"
  },
  "body": {
    "command": "ritman",
    "lastName": "cli",
    "name": "ritchie",
    "url": "https://ritchiecli.io/"
  }
}
```

2. You can easily use all the necessary HTTP Verbs for testing on the "method" field: GET, PUT, POST, PATCH, HEAD, OPTIONS, DELETE
3. you can also add an authorization token for private API testing in the "headers" field:

```json
{
  "Authorization": "<Bearer jwt> or <basic> or wathever token that fits your Authorization auth header"
}
```

4. The "body" field supports any valid JSON format, so you can fullfill with your disired test content.

## Command

```bash
rit http generate http-config
```

## Requirements

- [Golang installed](https://golang.org/doc/install)

## Demonstration

![Alt Text](https://media1.giphy.com/media/DvpJneeGh2v9TsHkQM/giphy.gif)
