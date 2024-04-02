# Mise-en-page cli formatter for Go

## Use

```go
redWord := mep.Color("word", mep.Red)
coloredLine := fmt.Sprintf("%s", redWord)
```
  

## Roadmap

- [x] Basic color and formatting (underline, strikethrough...)
- [ ] Components
  - [ ] Input
  - [ ] Select (radio) n == 1
  - [ ] Select (checkbox) n >= 1

> NOTE: v1

- [ ] Complex formatting (errors, warning, etc.)
- [ ] Choose output (stdout, buffer, etc.)
- [ ] Sanitize input

## Reference

[publish go package](https://go.dev/doc/modules/publishing)
