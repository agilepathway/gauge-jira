# Only one specs directory can be provided

tags: java, dotnet, ruby, python, js

Gauge documentation plugins allow for multiple specs directories to be passed in as command-line
arguments.  Our Jira plugin can only accept one though, as we match up the specs directory with
the specs Git URL that is also provided by the plugin user.

* Initialize an empty Gauge project

## The plugin fails if more than one specs directory is provided

* Publish Jira Documentation for two projects

* Console should contain "Aborting: this plugin only accepts one specs directory as a command-line argument."
