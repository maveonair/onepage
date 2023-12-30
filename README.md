# OnePage

OnePage is a simple web application that serves a single page, allowing users to edit content in Markdown format and view it rendered as HTML. It's designed to be lightweight and easy to use, perfect for situations where you need a single, easily editable web page.

Usage

1. Clone the Repository:
```bash
git clone https://github.com/maveonair/onepage.git
cd onepage
```

2. Build the Application:
```bash
make build
```

3. Run the application:
```bash
./dist/onepage
```

The application creates the file `page.md`, if it does not already exist, which contains the content of the website and is available at http://localhost:8080. The content can be edited there in Markdown format and is rendered as HTML.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.