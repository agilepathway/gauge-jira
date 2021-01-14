# Functional tests

The functional tests are themselves Gauge tests.

The scaffolding for the tests has been lifted and shifted from the
[gauge-tests](https://github.com/getgauge/gauge-tests) repository. When adding more Spectacle
functional tests definitely browse the `gauge-tests` repository for inspiration and ideas.

The functional tests run on every push and pull request, triggered by
[our functional test GitHub Action](../.github/workflows/functional-test.yml).

### Running the functional tests locally
- [Install Gauge](https://docs.gauge.org/getting_started/installing-gauge.html)

- [Install language plugin](https://docs.gauge.org/plugin.html) by running<br>
  ```
  gauge install java
  ```

- [Install Spectacle](../README.md#installation)
(you may want to [install from source](../README.md#build-from-source) to test your latest code)

- Clone this repo.

- Executing specs

  ```
  ./gradlew clean ft # On Linux or Mac
  gradlew.bat clean ft # On Windows
  ```

This will also compile all the supporting code implementations and run your specs in parallel.
