# Specs published to Jira have a link to the Git file

tags: java, dotnet, ruby, python, js

## When there is an HTTPS git config the Jira spec has a link to the Git file

* Initialize an empty Gauge project

* Add "HTTPS" Git config to project

* Create a scenario linked to Jira issue(s) "JIRAGAUGE-1"

* Publish Jira Documentation for the current project

* Console output should be "Published specifications to 1 Jira issue"

* Jira issue "JIRAGAUGE-1" description should contain basic scenario with Git link

## When there is an SSH git config the Jira spec has a link to the Git file

* Initialize an empty Gauge project

* Add "SSH" Git config to project

* Create a scenario linked to Jira issue(s) "JIRAGAUGE-1"

* Publish Jira Documentation for the current project

* Console output should be "Published specifications to 1 Jira issue"

* Jira issue "JIRAGAUGE-1" description should contain basic scenario with Git link

## When there is no git config the Jira spec does not have a link to the Git file

* Initialize an empty Gauge project without a Git config

* Create a scenario linked to Jira issue(s) "JIRAGAUGE-1"

* Publish Jira Documentation for the current project

* Console output should be "Published specifications to 1 Jira issue"

* Jira issue "JIRAGAUGE-1" description should contain basic scenario without Git link

## When in detached HEAD state the Jira spec does not have a link to the Git file

* Initialize an empty Gauge project

* Add "HTTPS" Git config to project

* Simulate Git detached HEAD

* Create a scenario linked to Jira issue(s) "JIRAGAUGE-1"

* Publish Jira Documentation for the current project

* Console output should be "Published specifications to 1 Jira issue"

* Jira issue "JIRAGAUGE-1" description should contain basic scenario without Git link
