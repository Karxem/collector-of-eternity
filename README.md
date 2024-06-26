[![Conventional Commits](https://github.com/Karxem/collector-of-eternity/actions/workflows/commit-lint.yml/badge.svg?branch=main)](https://github.com/Karxem/collector-of-eternity/actions/workflows/commit-lint.yml)
# Collector of Eternity - Top Down Game

Collector of Eternity is a small top down rts game that utilizes the ebitengine (formerly ebiten), a simple 2d game engine for go.

# Why? - Chosen Technologies

Why use ebitengine and go if there are big enignes like unity, ue5 or godot out there.

Bigger game engines simplify many things but also come with bloat and a specific approach.
Choosing a more light engine gives the opportunity to learn concepts from ground up and how games work.

Go got similiar benchmarks to compiled languages like C or C++, but has the simplicity of a garbage collected language.

# Composition over Inheritance
There are no classes in go and you can say, that there is no inheritance concept in either.
Go uses a concept called composition where a struct is used to form other objects.

## Running the Game

### Prerequisites

Before running the game, ensure you have the following installed:

- Go programming language

### Steps to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Karxem/collector-of-eternity.git
   cd collector-of-eternity
   ```

2. **Run the game:**

   ```bash
   go run .
   ```

   This command will build and start the game.

3. **Build for release:**

   ```bash
   go build -o ../dist/collector-of-eternity.exe
   ```

   This command will compile the project to the dist folder.

## Contributing

If you'd like to contribute to Collector of Eternity, please fork the repository and create a pull request.

**Notice:** Semantic commits are enforced with a pipeline. Only the following commit scopes are allowed: `feat`, `fix`, `chore`, `refactor`, `build`, `ci`, `docs`.

## License

This project is licensed under the [MIT License](LICENSE).
