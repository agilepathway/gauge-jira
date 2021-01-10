package com.thoughtworks.gauge.test.jira;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.http.HttpRequest.BodyPublishers;
import java.util.Base64;

import org.json.JSONObject;

public class IssueDescriptionAPIClient {

    public static String getDescriptionForIssue(String issueKey) {
        HttpResponse<String> rawJsonResponse = sendIssueRequest(getDescriptionRequest(issueKey));
        JSONObject jsonResponse = new JSONObject(rawJsonResponse.body());
        return jsonResponse.getJSONObject("fields").getString("description");
    }

    public static void setIssueDescription(String description, String issueKey) {
        JSONObject descriptionBody = new JSONObject().put("fields", new JSONObject().put("description", description));
        HttpRequest request = updateDescriptionRequest(issueKey, descriptionBody.toString());
        HttpResponse<String> response = sendIssueRequest(request);
        if (response.statusCode() != 204) {
            throw new IllegalStateException(
                    "Failed to set Jira description: " + response.body() + response.statusCode());
        }
    }

    private static HttpRequest getDescriptionRequest(String issueKey) {
        HttpRequest.Builder builder = baseDescriptionRequest();
        builder.uri(URI.create(descriptionAPIURL(issueKey) + "?fields=description"));
        return builder.build();
    }

    private static HttpRequest updateDescriptionRequest(String issueKey, String body) {
        HttpRequest.Builder builder = baseDescriptionRequest();
        builder.uri(URI.create(descriptionAPIURL(issueKey)));
        builder.PUT(BodyPublishers.ofString(body));
        return builder.build();
    }

    private static HttpResponse<String> sendIssueRequest(HttpRequest request) {
        HttpClient client = HttpClient.newBuilder().version(HttpClient.Version.HTTP_1_1).build();
        try {
            return client.send(request, HttpResponse.BodyHandlers.ofString());
        } catch (IOException | InterruptedException e) {
            throw new IllegalStateException("Exception when sending Jira issue request", e);
        }
    }

    private static HttpRequest.Builder baseDescriptionRequest() {
        String jiraUsername = System.getenv("JIRA_USERNAME");
        String jiraToken = System.getenv("JIRA_TOKEN");
        return HttpRequest.newBuilder()
            .header("Content-Type", "application/json")
            .header("Authorization", basicAuth(jiraUsername, jiraToken));
    }

    private static String descriptionAPIURL(String issueKey) {
        String jiraBaseUrl = System.getenv("JIRA_BASE_URL");
        return String.format("%1$s/rest/api/latest/issue/%2$s", jiraBaseUrl, issueKey);
    }

    private static String basicAuth(String username, String password) {
        return "Basic " + Base64.getEncoder().encodeToString((username + ":" + password).getBytes());
    }
    
}
