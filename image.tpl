pub const {{.Name}} = Image{
    .name = "{{.Name}}",
    .width = {{.Width}},
    .height = {{.Height}},
    .data = &.{ {{.BytesString}} },
};
