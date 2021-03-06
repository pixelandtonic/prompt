# Release Notes for Prompt

### Unreleased

### 1.1.8 - 2020-08-10

## Fixed
- Fixed an issue when using the confirm prompt.

### 1.1.7 - 2020-05-26

## Changed
- You can now override the global `AppendQuestionMarksOnAsk` setting and append a question mark on a per prompt basis using `AppendQuestionMark`.

### 1.1.6 - 2020-05-26

## Fixed
- Fixed a bug where the library was referencing the incorrect Go version.

### 1.1.5 - 2020-05-26

## Changed
- The `AppendQuestionMarksOnAsk` global option now defaults to `false`.

### 1.1.4 - 2020-05-26

## Fixed
- Fixed a nil pointer dereference.

### 1.1.3 - 2020-05-26

## Changed
- Validator is still ran even if input is empty.

## Fixed
- Corrected some behavior that would return default values when input was provided.

### 1.1.2 - 2020-05-23

## Changed
- Provide output that we are using the default when provided information is not valid.

### 1.1.1 - 2020-05-23

## Fixed
- Fixed an issue when validators were not running.
- Fixed an issue when selecting an option that was greater than the list. ([#3](https://github.com/pixelandtonic/prompt/issues/3))

### 1.1.0 - 2020-05-06

## Fixed
- Fixed an error where select would not account for zero-based indexes during selection. ([#2](https://github.com/pixelandtonic/prompt/issues/2))

### 1.0.0 - 2020-05-05
- Initial release.
