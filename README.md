# Pokedex CLI

Pokedex CLI is a command-line interface (CLI) application written in Go that allows users to explore the world of Pokémon. With this tool, you can navigate through various location areas, catch Pokémon, and manage your own Pokédex.

## Commands

Here are the available commands you can use in Pokedex CLI:

- `help`: Prints the Help Menu.
- `map`: Lists the next page of location areas.
- `mapb`: Lists the previous page of location areas.
- `explore {location_area}`: Lists the Pokémon in a specified location area.
- `catch {pokemon_name}`: Attempts to catch a Pokémon and add it to your Pokédex.
- `inspect {pokemon_name}`: View information about a caught Pokémon.
- `pokedex`: View all the Pokémon in your Pokédex.
- `exit`: Exits the Pokedex CLI application.

## Installation

To install Pokedex CLI, follow these steps:

1. Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. Clone the repository:

   ```bash
   git clone https://github.com/BoschLeith/pokedexcli.git
   ```

3. Navigate to the project directory:
  
    ```bash
    cd pokedexcli
    ```

4. Build the application:
  
    ```bash
    go build
    ```

5. Run the application:

    ```bash
    ./pokedexcli
    ```