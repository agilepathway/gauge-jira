# How to contribute

Firstly thanks for thinking of contributing - the project is [open source](https://opensource.guide/how-to-contribute/) and all contributions are very welcome :slightly_smiling_face: :boom: :thumbsup:

[How to report a bug or suggest a new feature](#how-to-report-a-bug-or-suggest-a-new-feature)

[How to make a contribution](#how-to-make-a-contribution)

[Local development](#local-development)
  * [Visual Studio Code](#visual-studio-code)
  * [Codespaces](#codespaces)
  * [Local development from scratch](#local-development-from-scratch)
    * [Dependencies](#dependencies)

[Unit tests](#unit-tests)
  * [Prerequisites for the unit tests](#prerequisites-for-the-unit-tests)
  * [Running the unit tests](#running-the-unit-tests)

[Functional tests](#functional-tests)

## How to report a bug or suggest a new feature

[Create an issue](../../issues), describing the bug or new feature in as much detail as you can.

## How to make a contribution

  * [Create an issue](../../issues) describing the change you are proposing.
  * [Create a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).  The project uses the _[fork and pull model](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/about-collaborative-development-models)_:
    * [Fork the project](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/working-with-forks)
    * Make your changes on your fork
        * [Update the tests or add new tests](./functional-tests/README.md) to cover the new behaviour.
    * Write a [good commit message(s)](https://chris.beams.io/posts/git-commit/) for your changes
    * [Create the pull request for your changes](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/proposing-changes-to-your-work-with-pull-requests)

### Visual Studio Code

The easiest way to set up your development environment (unless you have [Codespaces](#codespaces), which is even easier) is to use [Visual Studio Code](https://code.visualstudio.com/)'s [Remote Containers](https://code.visualstudio.com/docs/remote/containers) functionality:
  * [System requirements](https://code.visualstudio.com/docs/remote/containers#_system-requirements)
  * [Fork the project](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/working-with-forks) 
  * [Open the local project folder in a container](https://code.visualstudio.com/docs/remote/containers#_quick-start-open-an-existing-folder-in-a-container)
  * Everything will then be setup for you.  You'll be able to [run the tests](./functional-tests/README.md) locally.

### Codespaces

If you have access to [GitHub Codespaces](https://github.com/features/codespaces/) (which allows full remote
development from within your browser or VS Code) then all you need to do is 
[fork the project](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/working-with-forks) 
and open it in Codespaces - easy!

### Local development from scratch

#### Dependencies

* [Go](https://golang.org)
* [Java](https://www.java.com) version 15 or above (for [running the tests](./functional-tests/README.md))
* [Gauge](https://gauge.org)


### Unit tests

The unit tests are written in [Go](https://golang.org)

#### Prerequisites for the unit tests

Set `GAUGE_PROJECT_ROOT` as an environment variable with the full path of the project root directory,
e.g. `/home/someuser/workspace/gauge-jira`

#### Running the unit tests

`go test -v ./...`

### Functional tests

[Functional tests documentation](./functional-tests/README.md)
