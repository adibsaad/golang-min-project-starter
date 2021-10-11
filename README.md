# Golang Minimal Project Starter

This repository can be used as a starting point for new Golang projects.

## Features

- Viper/Cobra (for CLI and Configuration)
- Zerolog (logging)

## How-to-use

- The default name of the module is `my-app`, so you'll want to change all references to it in the code first.

- The settings configuration struct is located under `internal/settings`.
  It reads in values from Viper through the config file (if supplied)
  or environment variables, prefixed with the value of `EnvPrefix`.
  Customize the settings struct to fit your needs

## Credits

The code to set up the configuration struct was taken from the [MottainaiCI project](https://github.com/MottainaiCI/mottainai-server).
