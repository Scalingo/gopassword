# Go Password

Simple password generator in Go. Use `crypto/rand`

```go
// Passowrd of 20 characters
gopassword.Generate()

// Password of 42 characters
gopassword.Generate(42)
```

## Release a New Version

Bump new version number in `CHANGELOG.md`.

Commit, tag and create a new release:

```sh
git add CHANGELOG.md
git commit -m "Bump v1.0.0"
git tag v1.0.0
git push origin master
git push --tags
hub release create v1.0.0
```

The title of the release should be the version number and the text of the release is the same as the changelog.
