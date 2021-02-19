# ghlicense
Command-line client to quickly look up a repo license.

## Usage

1. Arm yourself with a [GitHub personal access
token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)

1. Build

        make

1. Run one of the commands (see `./ghlicense --help`)

        $ export GITHUB_ACCESS_TOKEN="1234567890abcdef1234567890abcde123456789"
        $ GITHUB_ACCESS_TOKEN=<token> ./ghlicense name https://github.com/petergardfjall/ghlicense
        MIT
        
        
