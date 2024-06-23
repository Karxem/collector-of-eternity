[![Conventional Commits](https://github.com/Karxem/pick-it-up/actions/workflows/commit-lint.yml/badge.svg?branch=main)](https://github.com/Karxem/pick-it-up/actions/workflows/commit-lint.yml)
# Pick it up - Top Down Game

Pik it up is a small top down rts game that utilizes the ebitengine (formerly ebiten), a simple 2d game engine for go.

# Why? - Chosen Technologies

Why use ebitengine and go if there are big enignes like unity, ue5 or godot out there.

Bigger game engines simplify many things but also come with bloat and a specific approach.
Choosing a more light engine gives the opportunity to learn concepts from ground up and how games work.

Go got similiar benchmarks to compiled languages like C or C++, but has the simplicity of a garbage collected language.

# Inheritance over Composition
I chose to not include a Entity Component System since part of this project is to get a deeper dive into the practice of inheritance.

## Running the Game

### Prerequisites

Before running the game, ensure you have the following installed:

- Go programming language

### Steps to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Karxem/pick-it-up.git
   cd pick-it-up
   ```

2. **Run the game:**

   ```bash
   go run .
   ```

   This command will build and start the game.

3. **Build for release:**

   ```bash
   go build -o ../dist/pick-it-up.exe
   ```

   This command will compile the project to the dist folder.

## Contributing

If you'd like to contribute to Pick it up, please fork the repository and create a pull request. 

## License

This project is licensed under the [MIT License](LICENSE).