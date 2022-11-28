---
id: mount
title: mount
hide_title: true
hide_table_of_contents: true
sidebar_label: mount
description: Mount an image
last_update:
  date: 2022-11-27T17:57:56-07:00
  author: blacktop
---
## ipsw idev img mount

Mount an image

```
ipsw idev img mount <image> <signature> [flags]
```

### Options

```
  -h, --help                help for mount
  -t, --image-type string   Image type to mount (default "Developer")
  -x, --xcode string        Path to Xcode.app (default "/Applications/Xcode.app")
```

### Options inherited from parent commands

```
      --color           colorize output
      --config string   config file (default is $HOME/.ipsw.yaml)
  -u, --udid string     Device UniqueDeviceID to connect to
  -V, --verbose         verbose output
```

### SEE ALSO

* [ipsw idev img](/docs/cli/ipsw/idev/img)	 - Image commands

