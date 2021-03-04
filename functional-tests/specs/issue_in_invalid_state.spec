# Jira issue with description that is in an invalid state

tags: java, dotnet, ruby, python, js

* Initialize an empty Gauge project

## Jira issue with description that is in an invalid state

Jira issues should only ever contain one Gauge examples section, currently (NB we may change this in
the future if we cater for more than one source repo separately publishing their Gauge specs to the
same Jira issue).

This test ensures that we prevent publishing of specs to a Jira issue which is in an invalid state
because it has more than one Gauge examples section in the Jira issue description.

A Jira issue could get into this invalid state either because of a manual edit, or because of a bug
(hypothetically) in the Gauge Jira plugin itself which inadvertently led to a duplicate examples
section instead of replacing the existing one.

* Set invalid description (with two Gauge sections) on Jira issue "JIRAGAUGE-10"

* Create a scenario linked to Jira issue(s) "JIRAGAUGE-10"

* Publish Jira Documentation for the current project

* Console should contain "JIRAGAUGE-10 is in an invalid state."
* Console should contain "It contains more than one Gauge examples section, but there should only ever be one or none."
* Console should contain "Remove all Gauge example sections from JIRAGAUGE-10 in Jira manually and then rerun the Gauge Jira plugin"
* Console should contain "No valid Jira specifications were found - so nothing to publish to Jira"

___

* Clear description on Jira issue "JIRAGAUGE-10"
