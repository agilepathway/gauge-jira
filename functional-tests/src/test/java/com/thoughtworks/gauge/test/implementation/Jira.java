package com.thoughtworks.gauge.test.implementation;

import static org.assertj.core.api.Assertions.assertThat;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Base64;

import com.thoughtworks.gauge.Step;

import org.json.JSONObject;

public class Jira {

    @Step("Jira issue <issuekey> description should contain basic scenario")
    public void verifyJiraIssueDescription(String issueKey) throws IOException, InterruptedException {
        String expectedScenario = expectedBasicScenario(issueKey);
        String issueDescription = getDescriptionForIssue(issueKey);
        assertThat(issueDescription).isEqualTo(expectedScenario);
    }

    private String expectedBasicScenario(String issueKey) {
        return """
            h1. %s

            tags:\040

            h2. Sample scenario

            * First step
            * Second step
            * Third step
            * Step with "two" "params"


            *_*

            """.formatted(issueKey);
    }

    private String getDescriptionForIssue(String issueKey) throws IOException, InterruptedException {
        String rawJsonResponse = sendJiraRequest(String.format("issue/%s?fields=description", issueKey));
        JSONObject jsonResponse = new JSONObject(rawJsonResponse);
        return jsonResponse.getJSONObject("fields").getString("description");
    }

    private String sendJiraRequest(String resourceAndQueryString) throws IOException, InterruptedException {
        String jiraBaseUrl = System.getenv("JIRA_BASE_URL");
        String jiraUsername = System.getenv("JIRA_USERNAME");
        String jiraToken = System.getenv("JIRA_TOKEN");
        String issueUrl = String.format("%1$s/rest/agile/latest/%2$s", jiraBaseUrl, resourceAndQueryString);
        HttpClient client = HttpClient.newBuilder().build();
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create(issueUrl))
            .header("Authorization", basicAuth(jiraUsername, jiraToken))
            .build();
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        return response.body();
    }

    private String basicAuth(String username, String password) {
        return "Basic " + Base64.getEncoder().encodeToString((username + ":" + password).getBytes());
    }
    
}
