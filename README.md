I can't figure out a cross-platform solution to expand relative path (possibly a symbolic link) to its absolute path.

So this utility.

```
$ normpath ../mytoken/solar.development.json

/Users/howard/p/qtum/qtumbook/examples/mytoken/solar.development.json
```

Install:

```
go get -u github.com/hayeah/normpath
```

See: https://unix.stackexchange.com/questions/36547/how-can-i-expand-a-relative-path-at-the-command-line-with-tab-completion
