# (Unofficial) appetize command line tools

[appetize.io](https://appetize.io/) is very useful, and I love it.
Make it more useful with this cli tool :)

## Config (optional)

You can put *default* parameters in `~/.appetize.yml`.

```
api_token: tok_hogehoge
platform: android
```

or set enviromental variables as below:

```
APPETIZE_API_TOKEN=tok_hogehoge
```

NOTE: `platform` and `api_token` can be overridden in executing the appetize command with `--platform` and `--token` parameters.

## Execute

### appetize upload

```
$ appetize upload app-debug.apk
```