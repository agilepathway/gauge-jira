Gauge-Jira
==========

[![Gauge Badge](https://gauge.org/Gauge_Badge.svg)](https://gauge.org)

[![build](https://github.com/agilepathway/gauge-jira/workflows/build/badge.svg)](https://github.com/agilepathway/gauge-jira/actions?query=workflow%3Abuild+event%3Apush+branch%3Amaster)
[![tests](https://github.com/agilepathway/gauge-jira/workflows/FTs/badge.svg)](https://github.com/agilepathway/gauge-jira/actions?query=workflow%3AFTs+event%3Apush+branch%3Amaster)
[![reviewdog](https://github.com/agilepathway/gauge-jira/workflows/reviewdog/badge.svg)](https://github.com/agilepathway/gauge-jira/actions?query=workflow%3Areviewdog+event%3Apush+branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/agilepathway/gauge-jira)](https://goreportcard.com/report/github.com/agilepathway/gauge-jira)

[![releases](https://img.shields.io/github/v/release/agilepathway/gauge-jira?color=blue&sort=semver)](https://github.com/agilepathway/gauge-jira/releases)
[![Supported Jira versions](https://img.shields.io/badge/supports-Jira%20Server%20%7C%20Jira%20Cloud-blue)](#supported-jira-versions)
[![License](https://img.shields.io/github/license/agilepathway/gauge-jira?color=blue)](LICENSE)


Publishes Gauge specifications to Jira. This is a plugin for [gauge](https://gauge.org/).
___
* [Why Publish Gauge Specs to Jira](#why-publish-gauge-specs-to-jira)
* [How to Use](#how-to-use)
  * [Typical Workflow](#typical-workflow)
  * [Supported Jira versions](#supported-jira-versions)
  * [Plugin setup](#plugin-setup)
  * [Running the plugin](#running-the-plugin)
  * [How to Link Gauge Specifications to Jira Issues](#how-to-link-gauge-specifications-to-jira-issues)
  * [Where in Jira are the Specs Published](#where-in-jira-are-the-specs-published)
  * [FAQs](#faqs)
* [Installation](#installation)
  * [Normal Installation](#normal-installation)
  * [Offline Installation](#offline-installation)
  * [Build from Source](#build-from-source)
* [Contributing](#contributing)
* [License](#license)

___


Why Publish Gauge Specs to Jira
-------------------------------

This plugin is aimed at teams who use Jira as an important part of their development lifecycle 
(having their [User Stories](https://www.mountaingoatsoftware.com/agile/user-stories) in Jira, for example).

It enables [living documentation](https://www.infoq.com/articles/book-review-living-documentation/) by publishing 
your Gauge specs right into the descriptions of the Jira stories that they relate to and therefore allowing 
everyone to see them, seamlessly.  

This is particularly useful if you are using [Specification by Example](http://specificationbyexample.com).

As [Gojko Adzic, the father of Specification by Example, says](https://gojko.net/2020/03/17/sbe-10-years.html#looking-forward-to-the-next-ten-years):

> *The big challenge related to tooling over the next 10 years will be in integrating better with Jira and its*
> *siblings. Somehow closing the loop so that teams that prefer to see information in task tracking tools get* 
> *the benefits of living documentation will be critical.*


How to Use
----------

### Typical Workflow

A typical workflow could be something like this:

1. collaborative story refinement sessions to come up with specification examples, using 
   [example mapping](https://cucumber.io/blog/bdd/example-mapping-introduction/) for instance
2. [write up the specification examples in Gauge](https://docs.gauge.org/writing-specifications.html)
   1. link specifications to individual Jira issues as required 
      ([see below](#how-to-link-gauge-specifications-to-jira-issues))
3. use this plugin in a [Continuous Integration (CI) pipeline](https://www.thoughtworks.com/continuous-integration)
   to publish (or republish) the specifications to Jira
4. [automate the specifications using Gauge](https://docs.gauge.org/writing-specifications.html#step-implementations) 
   whenever possible (not essential, there's still value even when not automated)
5. continue the cycle throughout the lifespan of the story: more conversations, more spec updates, 
   more automated publishing to Jira


### Supported Jira versions

The plugin supports both [Jira Server](https://www.atlassian.com/software/jira/latest-version?tab=server)
and [Jira Cloud](https://www.atlassian.com/software/jira/enterprise).

If you find a problem with a particular version of Jira Server or Jira Cloud, please
[raise an issue](../../issues)


### Plugin setup

There are three variables to configure, as either:

1. environment variables

2. properties in a 
   [properties file](https://docs.gauge.org/configuration.html#local-configuration-of-gauge-default-properties),
   e.g. `<project_root>/env/default/anythingyoulike.properties`

The four variables to configure are:

`JIRA_BASE_URL` e.g. `https://example.com` for Jira Server, or `https://example.atlassian.net` for Jira Cloud

`JIRA_USERNAME` e.g. `joe.bloggs` for Jira Server, or `joe@example.com` for Jira Cloud.
This user must have permissions to edit issues in the Jira projects that the specs will be linked to.

`JIRA_TOKEN` The Jira token is the password for the given `JIRA_USERNAME` if using Jira Server, or an
[api token](https://confluence.atlassian.com/cloud/api-tokens-938839638.html) if using Jira Cloud.


### Running the plugin (i.e. publishing specs to Jira)

`gauge docs jira`

or, if you want to specify a different directory to the default `specs` directory

`gauge docs jira <path to specs dir>`


### How to Link Gauge Specifications to Jira Issues

Simply add one or more 
[Jira issue keys](https://support.atlassian.com/jira-software-cloud/docs/what-is-an-issue/#:~:text=Issue%20keys%20are%20unique%20identifiers),
anywhere in the Gauge specification.

* You can link more than one Jira issue to a Gauge specification.  Just add the Jira issue keys
  anywhere in the Gauge spec (they don't all need to be together, although it may make sense from
  a readability point of view to keep them together)

* You can also link more than one Gauge specification to the same Jira issue(s).

* If you link to a Jira issue multiple times from the same Gauge specification, the specification
  will only be published to Jira once (rather than appearing in duplicate in Jira).

* Linked Jira issues always pertain to the entire Gauge specification - there is no current ability
  to publish at the individual scenario level of granularity.

* Jira issue keys are of the following format:

`<projectkey>-<issuenumber>`

e.g.

`MYPROJECT-1`

* You can specify the Jira issue keys in a normal 
  [Gauge comment](https://docs.gauge.org/writing-specifications.html#how-to-add-comments-in-a-specification),
  (i.e. in normal plain text anywhere in the specification), e.g.

  `Linked Jira issue: MYPROJECT-1` or `Linked Jira issues: MYPROJECT-1, MYPROJECT-2`

* Or you could specify them as url links, e.g.

  `https://example.com/browse/MYPROJECT-1, https://example.com/browse/MYPROJECT-2`

  or

  `[MYPROJECT-1](https://example.com/browse/MYPROJECT-1), [MYPROJECT-2](https://example.com/browse/MYPROJECT-2)`

* Or you can specify the Jira issue keys as [Gauge tags](https://docs.gauge.org/writing-specifications.html#tags)
  if you wish, e.g.

  `Tags: MYPROJECT-1, MYPROJECT-2`


### Where in Jira are the Specs Published

The specifications are appended to the existing contents of the Jira description field.

You can (and should) republish at will, as the plugin replaces any previously published specs 
with the latest version.  Recommended approach is to republish them to Jira automatically on
every change to the specs, as part of a Continuous Integration pipeline.

We chose to publish to the description field because it is a core field in Jira, always
present and clearly visible in all issue pages in Jira.


### FAQs

1. Can the specifications be edited in Jira and synced back into the Gauge specs?

   No.  Only make edits in the Gauge specifications themselves.  We include a message in
   Jira warning not to make edits to the specifications in Jira.

2. Is it safe to publish the specs to Jira multiple times?

   Yes.  The plugin replaces any previously published specs with the latest version.

3. What happens if the Gauge specs are longer than the Jira description field maximum length?

   The default maximum length of Jira fields is 32,767 characters.  This is long enough for
   most cases.  If the specs for a given issue exceed the maximum length, the plugin will skip
   publishing that issue and report this in the console output.  The recommended solution if
   that happens is to get an administrator of your Jira instance (or the Jira support team, if
   you are using Jira Cloud) to 
   [increase the maximum field length](https://confluence.atlassian.com/adminjiraserver073/configuring-advanced-settings-861253966.html).



Installation
------------


### Normal Installation

```
gauge install jira
```
To install a specific version of jira plugin use the ``--version`` flag.

```
gauge install jira --version $VERSION
```


### Offline Installation

Download the plugin zip from the [Github Releases](https://github.com/agilepathway/gauge-jira/releases),
or alternatively (if you want to experiment with an unreleased version, which is not recommended) from the
[artifacts](https://docs.github.com/actions/managing-workflow-runs/downloading-workflow-artifacts) in the
[`Store distros`](../../actions?query=workflow%3A%22Store+distros%22) GitHub Action (NB you must be logged
in to GitHub to be able to retrive the artifacts from there).

use the ``--file`` or ``-f`` flag to install the plugin from  zip file.

```
gauge install jira --file ZIP_FILE_PATH
```

### Build from Source


#### Requirements
* [Golang](http://golang.org/)


#### Compiling

```
go run build/make.go
```

For cross-platform compilation

```
go run build/make.go --all-platforms
```


#### Installing
After compilation

```
go run build/make.go --install
```


#### Creating distributable

Note: Run after compiling

```
go run build/make.go --distro
```

For distributable across platforms: Windows and Linux for both x86 and x86_64

```
go run build/make.go --distro --all-platforms
```


Contributing
------------

See the [CONTRIBUTING.md](./CONTRIBUTING.md)


License
-------

`Gauge-Jira` is released under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text.
