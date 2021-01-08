Jira spec with scenarios
===============================

tags: jira, java, dotnet, ruby, python, js

* Initialize a project named "spec_with_scenarios" without example spec

Basic spec with one scenario
-------------------------------------

* Create a scenario "Sample scenario" in specification "Basic spec execution" with the following steps with implementation 

   |step text               |implementation                                          |
   |------------------------|--------------------------------------------------------|
   |First step              |"inside first step"                                     |
   |Second step             |"inside second step"                                    |
   |Third step              |"inside third step"                                     |
   |Step with "two" "params"|"inside step with parameters : " + param0 + " " + param1|

* Publish Jira Documentation for the current project

* Console should contain "Successfully exported specs to Jira"
