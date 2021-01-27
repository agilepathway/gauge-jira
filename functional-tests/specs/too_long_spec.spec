Jira spec that exceeds the maximum length of the description field
==================================================================

tags: java, dotnet, ruby, python, js

* Initialize an empty Gauge project

Jira spec that exceeds the maximum length of the description field
------------------------------------------------------------------

* Create a scenario exceeding maximum default Jira field length linked to Jira issue(s) "JIRAGAUGE-9"

* Publish Jira Documentation for the current project

* Console should contain "The specification(s) for issue JIRAGAUGE-9 exceeds the default Jira maximum field length of 32767 characters."

* Console should contain "You can ask your Jira administrator to increase the maximum field length (or raise a support ticket if you are on Jira Cloud)."

* Console should contain "Failed to publish issue JIRAGAUGE-9"

* Console should contain "No valid Jira specifications were found - so nothing to publish to Jira"
