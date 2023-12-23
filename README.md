# tmpl
> Render go templates from cli

## Usage

**vars.json**
```json
{ "one": 1 }
```

**text.tpl**
```
You are number {{ .one }}!
```

```shell
$ tmpl -in text.tpl -data vars.json
You are number 1!
```