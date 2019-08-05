# 🐋 Baleia – A template engine used to manage Docker images templates repositories

Baleia is a template engine used to manage Docker images templates repositories.
The project is based on three philosophies *KISS*, *DRY* and *YAGNI*.

---

## Installation

K9s is available on Linux, OSX and Windows platforms.

* Binaries for Mac OS, Linux and Windows are available as tarballs in the [release](https://github.com/julienbreux/baleia/releases) page.

* Via Homebrew (Mac OS) or LinuxBrew (Linux)

   ```shell
   brew tap julienbreux/baleia
   brew install baleia
   ```

* Building from source
   Baleia was built using go 1.12 or above. In order to build Baleia from source you must:
   1. Clone this repository
   2. Enable Go module via env var `GO111MODULE=on`
   3. Add the following command in your go.mod file

      ```text
      replace (
        github.com/julienbreux/baleia => CLONED_GIT_REPOSITORY
      )
      ```

   4. Build and run the executable

        ```shell
        go run main.go
        ```

   5. Use it

        ```shell
        ./baleia
        ```

---

## Contact Info

1. **Email**:   julien.breux@gmail.com
2. **GitHub**:  [@JulienBreux](https://github.com/JulienBreux)
3. **Twitter**: [@JulienBreux](https://twitter.com/JulienBreux)

---

## Security info

### GPG Signature

You can download Julien Breux's public key to verify the signature.

```shell
gpg --keyserver hkps://hkps.pool.sks-keyservers.net --recv-keys 0BD023FA
```
