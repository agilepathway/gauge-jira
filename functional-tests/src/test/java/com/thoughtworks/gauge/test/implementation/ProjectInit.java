package com.thoughtworks.gauge.test.implementation;

import com.thoughtworks.gauge.Step;
import com.thoughtworks.gauge.test.common.GaugeProject;
import com.thoughtworks.gauge.test.common.Util;
import com.thoughtworks.gauge.test.common.builders.ProjectBuilder;

public class ProjectInit {

    private ThreadLocal<GaugeProject> currentProject = new ThreadLocal<GaugeProject>();

    @Step("Initialize an empty Gauge project")
    public void projectInit() throws Exception {
        currentProject.set(new ProjectBuilder()
                .withLangauge(Util.getCurrentLanguage())
                .withProjectName("gauge_jira_specs")
                .withoutExampleSpec()
                .build(false));
    }
}
