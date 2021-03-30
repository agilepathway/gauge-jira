package com.thoughtworks.gauge.test.jira;

import static org.assertj.core.api.Assertions.assertThat;
import static com.thoughtworks.gauge.test.jira.IssueDescriptionAPIClient.getDescriptionForIssue;
import static com.thoughtworks.gauge.test.jira.IssueDescriptionAPIClient.setIssueDescription;

import com.thoughtworks.gauge.Step;

public class Jira {

    @Step("Jira issue <issuekey> description should contain basic scenario named <scenario name>")
    public void verifyJiraIssueDescriptionForBasicScenario(String issueKey, String scenarioName) {
        String expectedScenario = expectedExamplesHeader() + expectedBasicScenarioHeader(scenarioName)
                + expectedBasicSpec() + expectedBasicScenarioFooter() + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step("Jira issue <issuekey> description should contain <originalDescription> and basic scenario")
    public void verifyJiraIssueContainsOriginalDescriptionAndBasicScenario(String issueKey,
            String originalDescription) {
        String expectedScenario = expectedOriginalDescription(originalDescription) + expectedExamplesHeader()
                + expectedBasicScenarioHeader(issueKey) + expectedBasicSpec() + expectedBasicScenarioFooter()
                + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step({ "Jira issue <issuekey> description should contain basic scenario",
            "Jira issue <issuekey> description should contain basic scenario with Git link" })
    public void verifyJiraIssueDescriptionForBasicScenario(String issueKey) {
        verifyJiraIssueDescriptionForBasicScenario(issueKey, issueKey);
    }

    @Step("Jira issue <issuekey> description should contain basic scenario without Git link")
    public void verifyJiraIssueDescriptionForBasicScenarioWithoutGitLink(String issueKey) {
        String expectedScenario = expectedExamplesHeader() + expectedBasicScenarioHeaderWithoutGitLink(issueKey)
                + expectedBasicSpec() + expectedBasicScenarioFooter() + expectedExamplesFooter();
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    @Step("Jira issue <issuekey> description should contain basic scenario, twice")
    public void verifyJiraIssueDescriptionForTwoBasicScenarios(String issueKey) {
        String expectedScenario = expectedExamplesHeader() + expectedBasicScenarioHeader(issueKey) + expectedBasicSpec()
                + expectedBasicSpec() + expectedBasicScenarioFooter() + expectedExamplesFooter();
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

    @Step("Set invalid description (with two Gauge sections) on Jira issue <issuekey>")
    public void setInvalidDescriptionWithTwoGaugeSections(String issueKey) {
        String exampleSection = expectedExamplesHeader() + expectedBasicScenarioHeader(issueKey) + expectedBasicSpec()
                + expectedBasicScenarioFooter() + expectedExamplesFooter();
        String invalidDescriptionWithTwoGaugeSections = exampleSection + exampleSection;
        setIssueDescription(invalidDescriptionWithTwoGaugeSections, issueKey);
    }

    private String expectedExamplesHeader() {
        return """

                ----
                ----
                h2.Specification Examples
                h3.Edit these examples in Git (link is below), not here in Jira
                """;
    }

    private String expectedExamplesFooter() {
        return """

                ----
                End of specification examples
                ----
                ----
                """;
    }

    private String expectedOriginalDescription(String originalDescription) {
        return """
                %s
                """.formatted(originalDescription);
    }

    private String expectedBasicScenarioHeader(String scenarioName) {
        /*
         * The expected URL is as per {@link
         * com.thoughtworks.gauge.test.git.Config.GitConfig#remoteOriginURL(String)}
         */
        String expectedGithubBaseURL = "https://github.com/example-user/example-repo/blob/master/specs";
        return """
                ----
                h3. %2$s
                [View or edit this spec in Git|%1$s/%2$s.spec]


                tags:\040

                """.formatted(expectedGithubBaseURL, scenarioName);
    }

    private String expectedBasicScenarioHeaderWithoutGitLink(String scenarioName) {
        return """
                ----
                h3. %s

                tags:\040

                """.formatted(scenarioName);
    }

    private String expectedBasicSpec() {
        return """
                h4. Sample scenario

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
