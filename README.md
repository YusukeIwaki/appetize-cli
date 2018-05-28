[![CircleCI](https://circleci.com/gh/YusukeIwaki/appetize-cli.svg?style=svg)](https://circleci.com/gh/YusukeIwaki/appetize-cli)

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

### NOTE

* `platform` and `api_token` can be overridden in executing the appetize command with `--platform` and `--token` parameters.
* appetize.yml can be easily generated with `appetize configure` command :)


## Execute

### appetize upload

```
$ appetize upload app-debug.apk

PublicKey:      1521kwdbrewp8
Created:        2018-05-23 01:27:55.617 +0000 UTC
Updated:        2018-05-23 01:27:55.617 +0000 UTC
Platform:       android
VersionCode:    1
PublicUrl:      https://appetize.io/app/1521kwdbrewp8
AppUrl: https://appetize.io/app/1521kwdbrewp8
ManageUrl:      https://appetize.io/manage/private_tkmvwt9tf5ag9jc

$ appetize upload 1521kwdbrewp8 app-debug.apk

PublicKey:      1521kwdbrewp8
Created:        2018-05-23 01:27:55.617 +0000 UTC
Updated:        2018-05-23 16:24:45.792 +0000 UTC
Platform:       android
VersionCode:    2
PublicUrl:      https://appetize.io/app/1521kwdbrewp8
AppUrl: https://appetize.io/app/1521kwdbrewp8
ManageUrl:      https://appetize.io/manage/private_tkmvwt9tf5ag9jc
```

### appetize create

```
$ appetize create https://example.com/files/app-debug.apk

PublicKey:      1521kwdbrewp8
Created:        2018-05-23 01:27:55.617 +0000 UTC
Updated:        2018-05-23 01:27:55.617 +0000 UTC
Platform:       android
Note:
LaunchUrl:
VersionCode:    1
PublicUrl:      https://appetize.io/app/1521kwdbrewp8
AppUrl: https://appetize.io/app/1521kwdbrewp8
ManageUrl:      https://appetize.io/manage/private_tkmvwt9tf5ag9jc
```

### appetize list

```
$ appetize list

data:   1521kwdbrewp8      2018-05-23 01:27:55.617 +0000 UTC
data:   ymnuvqm545gnc      2018-05-23 01:24:57.485 +0000 UTC
data:   byxn1f9mpgpwg      2017-10-12 13:23:11.871 +0000 UTC
```

### appetize show

```
$ appetize show 1521kwdbrewp8

PublicKey:      1521kwdbrewp8
Created:        2018-05-23 01:27:55.617 +0000 UTC
Updated:        2018-05-23 01:27:55.617 +0000 UTC
Disabled:       false
Platform:       android
VersionCode:    1
Bundle: com.example.helloworld
Name:   Hello World
Note:
AppVersionName: 1.0
AppVersionCode: 1
IconUrl:        https://s3.amazonaws.com/appetizeio-static/icons/uj027erw8z0q8rgbqutpj82j54_icon.png
ViewUrl:        https://appetize.io/app/1521kwdbrewp8
```

### appetize update

```
$ appetize update 1521kwdbrewp8 --disabled=true --note=hogehoge --timeout=30

PublicKey:      1521kwdbrewp8
Created:        2018-05-23 01:27:55.617 +0000 UTC
Updated:        2018-05-23 16:24:45.792 +0000 UTC
Disabled:       true
Timeout:        30
Platform:       android
VersionCode:    2
Bundle: com.example.helloworld
Name:   Hello World
Note:   hogehoge
AppVersionName: 1.0
AppVersionCode: 1
IconUrl:        https://s3.amazonaws.com/appetizeio-static/icons/uj027erw8z0q8rgbqutpj82j54_icon.png
ViewUrl:        https://appetize.io/app/1521kwdbrewp8
```

Note that we have to always set `--disabled=true` for keeping an app disabled. 
Appetize API is a little buggy :(

### appetize delete

```
$ appetize delete 1521kwdbrewp8 zf98sdbrewp89 byxn1f9mpgpwg

1521kwdbrewp8   OK
zf98sdbrewp89   error:App not found
byxn1f9mpgpwg   OK
```
