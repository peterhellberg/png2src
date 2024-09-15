# png2src

Convert `.png` into source code as described by the template used.

By default this tool expects a `image.tpl` in the current working directory.

## Example template

When I want a Zig :zap: constant of type `Image` I might have a template such as this:

```zig
pub const {{.Name}} = Image{
    .name = "{{.Name}}",
    .width = {{.Width}},
    .height = {{.Height}},
    .data = &.{ {{.BytesString}} },
};
```
