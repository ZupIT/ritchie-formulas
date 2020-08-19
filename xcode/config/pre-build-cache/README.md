# Ritchie Formula

## command

```bash
rit xcode config pre-build-cache
```

## description

Many times XCode users require to clean the cache to reload some project components,
leading to a build time of 3min~5min. This pre-build cache cleanup only affect
specific project files and prevent long project rebuilds while always
keeping new changes visible.
