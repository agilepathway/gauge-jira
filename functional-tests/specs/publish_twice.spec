Publish one Jira spec twice
===============================

tags: java, dotnet, ruby, python, js

* Initialize a project named "spec_with_scenarios" without example spec

Existing specs already in a Jira issue are overwritten when publishing
----------------------------------------------------------------------

* Create a basic scenario linked to Jira issue(s) "JIRAGAUGE-7"

* Publish Jira Documentation for the current project

* Console should contain "Published specifications to 1 Jira issue"

* Publish Jira Documentation for the current project

* Console should contain "Published specifications to 1 Jira issue"

* Jira issue "JIRAGAUGE-7" description should contain basic scenario

___

* Set description "" on Jira issue "JIRAGAUGE-7"
