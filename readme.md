## TrackSkip

**TrackSkip** is a utility that extracts and redirects to the real URL embedded within tracking URLs. It offers a library, a command-line interface, and a web server to help users skip past those pesky tracking links.

### Features

- **Library**: Integrate `trackskip` functionalities into your Go project.
- **CLI**: Quickly get the real URL right from your terminal.
- **Server**: Host your own redirect service to bypass tracking URLs.

### Installation

To get started with `trackskip`, clone this repository:

```bash
git clone https://github.com/PatrickTulskie/trackskip.git
cd trackskip
```

Then, build the project:

```bash
go build .
```

### Usage

1. **Library**:

   You can import `trackskip` in your Go project as follows:

   ```go
   import "github.com/PatrickTulskie/trackskip/trackskip"
   ```

   And use the `ExtractRedirectURL` function to retrieve the real URL.

2. **CLI**:

   After building, you can use the command-line tool as follows:

   ```bash
   ./trackskip "https://www.example.com/?q=https%3A%2F%2Fgoogle.com"
   ```

   Or simply:

   ```bash
   echo "https://www.example.com/?q=https%3A%2F%2Fgoogle.com" | ./trackskip
   ```

3. **Server**:

   Run the server with:

   ```bash
   go run server.go
   ```

   Once running, you can visit `http://localhost:8080/?q=<TRACKING_URL>` to get redirected to the real URL.

### Docker

If you're running docker, you can use it like so:

1. CLI

```
URL="<TRACKING_URL>" docker-compose run trackskip-cli
```

2. Server

```
docker-compose up trackskip-server
```

### Testing

Tests are available for the library, CLI, and server. Run them using:

```bash
go test ./...
```

### Contributions

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to change.

### License

[MIT](https://choosealicense.com/licenses/mit/)
