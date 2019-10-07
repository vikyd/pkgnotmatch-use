# pkgnotmatch-use

Test if package path last element not match package name

# Usage

Download or git clone this respo directly.

Run:

```sh
go run main.go
```

# Conclusion

Golang can use package which its name is not match its path.

The customer should use the package name instead of last element of its path.

- can be: `abc.F01()`
- can not be: `pkgnotmatch.F01()`

About import:

- can be: `import "github.com/vikyd/pkgnotmatch"`
  - used as: `abc.F01()`
- can also be: `import abc "github.com/vikyd/pkgnotmatch"`
  - used as: `abc.F01()`
- can also be: `import othername "github.com/vikyd/pkgnotmatch"`
  - used as: `othername.F01()`
