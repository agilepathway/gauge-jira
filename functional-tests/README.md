# Functional tests

The functional tests are themselves Gauge tests.

They provide full end-to-end test coverage by using the
plugin to publish Gauge specs to an actual Jira instance.

The scaffolding for the tests has been lifted and shifted from the
[gauge-tests](https://github.com/getgauge/gauge-tests) repository. When adding more
functional tests definitely browse the `gauge-tests` repository for inspiration and ideas.

The functional tests run on every push and pull request, triggered by
[our functional test GitHub Action](../.github/workflows/functional-test.yml),
and running against both an actual Jira Server instance and an actual Jira Cloud instance.

### Running the functional tests locally

- Prerequisites

  * Java (version 15 or above)

  * [Install Gauge](https://docs.gauge.org/getting_started/installing-gauge.html)

  * [Install the Java language plugin](https://docs.gauge.org/plugin.html): `gauge install java`

  * [Install the Jira plugin](../README.md#installation)
(you may want to [install from source](../README.md#build-from-source) to test your latest code)

  * [Setup the Jira plugin](../README.md#plugin-setup) - ask a [core maintainer](../.github/CODEOWNERS) for the values
  to populate the required environment variables (or properties file) with.

- Running the tests

  ```
  ./gradlew clean ft # On Linux or Mac
  
  gradlew.bat clean ft # On Windows
  ```
