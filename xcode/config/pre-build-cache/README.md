# Description

Many times XCode users require to clean the cache to reload some project components,
leading to a build time of 3min~5min. This pre-build cache cleanup only affect
specific project files and prevent long project rebuilds while always
keeping new changes visible.

## Command

```bash
rit xcode config pre-build-cache
```

## Requirements

- [Node](https://nodejs.org/en/)
- [Npm](https://www.npmjs.com/get-npm)

## Demonstration

![gif](https://github.com/ZupIT/ritchie-formulas/raw/master/xcode/config/pre-build-cache/docs/xcode-pre-build-cache.gif)
