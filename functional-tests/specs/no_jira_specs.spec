Non-Jira spec with one scenario
===============================

tags: java, dotnet, ruby, python, js

* Initialize an empty Gauge project

Basic spec
-------------------------------------

* Create a scenario "Sample scenario" in specification "Spec not linked to Jira" with the following steps with implementation 

   |step text               |implementation                                          |
   |------------------------|--------------------------------------------------------|
   |First step              |"inside first step"                                     |
   |Second step             |"inside second step"                                    |
   |Third step              |"inside third step"                                     |
   |Step with "two" "params"|"inside step with parameters : " + param0 + " " + param1|

* Publish Jira Documentation for the current project

* Console output should be "No valid Jira specifications were found - so nothing to publish to Jira"
