package com.thoughtworks.gauge.test.implementation;

import static com.thoughtworks.gauge.test.common.GaugeProject.getCurrentProject;

import com.thoughtworks.gauge.Step;
import com.thoughtworks.gauge.test.common.ExecutionSummary;
import com.thoughtworks.gauge.test.common.ExecutionSummaryAssert;

public class Execution {

    private String getFormattedProcessOutput() {
        return "\n*************** Process output start************\n" + getCurrentProject().getLastProcessStdout()
                + "\n*************** Process output end************\n";
    }

    private ExecutionSummaryAssert assertOn(ExecutionSummary summary, boolean result) {
        return ExecutionSummaryAssert.assertThat(summary).withFailMessage(getFormattedProcessOutput())
                .hasSuccess(result);
    }

    @Step("Publish Jira Documentation for the current project")
    public void publishJiraDocumentationForCurrentProject() throws Exception {
        assertOn(getCurrentProject().publishJiraDocumentation(), true);
    }

    @Step("Publish Jira Documentation for the current project with no <variable> configured")
    public void publishJiraDocumentationForCurrentProjectWithConfigVarUnset(String configVar) throws Exception {
        assertOn(getCurrentProject().publishJiraDocumentationWithConfigVarUnset(configVar), false);
    }

    @Step("Publish Jira Documentation for two projects")
    public void publishJiraDocumentationForTwoProjects() throws Exception {
        assertOn(getCurrentProject().publishJiraDocumentationForTwoProjects(), true);
    }
}
