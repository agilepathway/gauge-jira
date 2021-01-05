package com.thoughtworks.gauge.test.implementation;

import static org.assertj.core.api.Assertions.assertThat;
import static com.thoughtworks.gauge.test.jira.IssueDescriptionAPIClient.getDescriptionForIssue;
import static com.thoughtworks.gauge.test.jira.IssueDescriptionAPIClient.setIssueDescription;

import com.thoughtworks.gauge.Step;

public class Jira {

    @Step("Jira issue <issuekey> description should contain basic scenario named <scenario name>")
    public void verifyJiraIssueDescriptionForBasicScenario(String issueKey, String scenarioName) {
        String expectedScenario = expectedExamplesHeader() + expectedBasicScenarioHeader(scenarioName) 
            + expectedBasicSpec() +  expectedBasicScenarioFooter() + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step("Jira issue <issuekey> description should contain <originalDescription> and basic scenario")
    public void verifyJiraIssueContainsOriginalDescriptionAndBasicScenario(String issueKey, String originalDescription) {
        String expectedScenario = expectedOriginalDescription(originalDescription) + expectedExamplesHeader()
            + expectedBasicScenarioHeader(issueKey) + expectedBasicSpec() + expectedBasicScenarioFooter() 
            + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step("Jira issue <issuekey> description should contain basic scenario")
    public void verifyJiraIssueDescriptionForBasicScenario(String issueKey) {
        verifyJiraIssueDescriptionForBasicScenario(issueKey, issueKey);
    }

    @Step("Jira issue <issuekey> description should contain basic scenario, twice")
    public void verifyJiraIssueDescriptionForTwoBasicScenarios(String issueKey) {
        String expectedScenario = expectedExamplesHeader() + expectedBasicScenarioHeader(issueKey)
            + expectedBasicSpec() + expectedBasicSpec()
            + expectedBasicScenarioFooter() + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step("Jira issue <issuekey> description should contain basic scenarios <scenario1>, <scenario2>")
    public void verifyJiraIssueDescriptionForTwoBasicScenarios(String issueKey, String scenario1, String scenario2) {
        String expectedDescription = expectedExamplesHeader() + expectedBasicScenarioHeader(scenario1) 
            + expectedBasicSpec() + expectedBasicScenarioFooter() + expectedBasicScenarioHeader(scenario2) 
            + expectedBasicSpec() + expectedBasicScenarioFooter() + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedDescription);
    }

    @Step("Set description <description> on Jira issue <issuekey>")
    public void setJiraIssueDescription(String description, String issueKey) {
        setIssueDescription(description, issueKey);
    }

    private String expectedExamplesHeader() {
        return """
        ----
        h2.Specification Examples
        """;
    }

    private String expectedExamplesFooter() {
        return """
            ------------------------------
            End of specification examples
            ----
            """;
    }

    private String expectedOriginalDescription(String originalDescription) {
        return """
            %s
            """.formatted(originalDescription);
    }

    private String expectedBasicScenarioHeader(String scenarioName) {
        return """
            h1. %s

            tags:\040

            """.formatted(scenarioName);
    }

    private String expectedBasicSpec() {
        return """
            h2. Sample scenario

            * First step
            * Second step
            * Third step
            * Step with "two" "params"

            """;
    }

    private String expectedBasicScenarioFooter() {
        return """

            *_*

            """;
    }
    
}
