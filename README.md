# Git Branch Puller

This Go program automates the process of pulling all local branches in a Git repository if their corresponding remote branches exist. It traverses a specified root directory, identifies Git repositories by the presence of a `.gitignore` file, and updates the branches in each repository.

## Features

- Recursively traverse directories to find Git repositories.
- Pull updates for all local branches that have corresponding remote branches.

## Requirements

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    ```
2. Navigate into the project directory:
    ```sh
    cd git-branch-puller
    ```
3. Build the Go application:
    ```sh
    go build -o git-branch-puller
    ```

## Usage

1. Run the built executable:
    ```sh
    ./git-branch-puller
    ```
2. Enter the root directory when prompted:
    ```sh
    Bitte geben Sie das Root-Verzeichnis ein: /path/to/root
    ```
3. The program will recursively traverse the specified root directory, find Git repositories, and pull updates for all local branches that have corresponding remote branches.

## Example

```sh
$ ./git-branch-puller
Bitte geben Sie das Root-Verzeichnis ein: /home/user/projects
Git-ähnliches Verzeichnis gefunden: /home/user/projects/repo1
Branch main in /home/user/projects/repo1 erfolgreich gepullt.
Branch dev in /home/user/projects/repo1 erfolgreich gepullt.
Git-ähnliches Verzeichnis gefunden: /home/user/projects/repo2
Branch main in /home/user/projects/repo2 erfolgreich gepullt.
Remote-Branch feature does not exist, skipping pull in /home/user/projects/repo2.
```

## How It Works

1. **Root Directory Input:** The program prompts the user to input a root directory.
2. **Directory Traversal:** It recursively traverses the given directory using `filepath.Walk`.
3. **Git Repository Detection:** For each directory, it checks for the presence of a `.gitignore` file to identify Git repositories.
4. **Branch Pulling:**
    - Retrieves all local branches using `git branch`.
    - Checks if the corresponding remote branch exists using `git ls-remote`.
    - For each local branch with a corresponding remote branch:
        - Checks out the branch.
        - Pulls updates from the remote branch.

## Error Handling

- The program outputs error messages if any issues occur during directory traversal, branch checking, or branch pulling.
- It skips directories that do not contain Git repositories and continues with the next ones.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix:
    ```sh
    git checkout -b feature-name
    ```
3. Commit your changes:
    ```sh
    git commit -m "Add new feature"
    ```
4. Push to the branch:
    ```sh
    git push origin feature-name
    ```
5. Create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

This project was inspired by the need to efficiently manage multiple Git repositories and keep local branches updated.
