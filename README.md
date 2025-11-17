# Go Web Template

A modern web application template using Go, Templ, and Tailwind CSS.

## Features

-   **Go** with Chi router
-   **Templ** for type-safe HTML templates
-   **Tailwind CSS v4** for styling
-   **TypeScript** support with Vite
-   **HTMX** for interactive UI
-   Hot reload for development

## Getting Started

### Prerequisites

-   Go 1.24+
-   Node.js and npm
-   Templ CLI

### Development

1. Install dependencies:

    ```bash
    npm install
    go install github.com/a-h/templ/cmd/templ@latest
    go mod tidy
    ```

1. Run the development server:

    ```bash
    # Run all watch tasks (CSS, TypeScript, Vite, and Templ with hot reload):
    npm run dev
    ```

    This will:

    - Watch for CSS changes and rebuild Tailwind
    - Watch for TypeScript changes
    - Watch for Vite build changes
    - Watch for Templ template changes and run the Go server with hot reload

1. Visit `http://localhost:7331` in your browser (the Templ proxy port for hot reload).

    **Note:** The Templ proxy runs on port 7331 and forwards requests to your Go server on port 8080. The proxy injects hot reload scripts into your HTML, so you should connect to port 7331 during development, not port 8080 directly.

## Project Structure

- `main.go` - Main application entry point (includes handlers)
- `*.templ` - Templ template files
- `utils/` - Templ Utility functions for Tailwind
- `css/input.css` - Tailwind CSS input file
- `dist/` - Built assets (CSS, JS) - served by the Go server
- `ts/` - TypeScript source files
- `static/` - Public directory

## Notes

- Organize your `handlers`/`routes`/`templ`/`pages`/`components`/`views`/`models` however you want.
- The generated `*.templ.go` files aren't commited.
- Change the `go.mod` module name to your desired name.
- You can change the name of your public directory in `vite.config.ts`, default is `static`.

## Scripts

- `npm run dev` - Run all watch tasks (CSS, TypeScript, Vite, Templ)
- `npm run build` - Build production assets
- `npm run templ:generate` - Generate templ Go files
- `npm run watch:css` - Watch and build CSS to dist directory
- `npm run lint` - Lint TypeScript files
- `npm run test` - Run tests
