package com.thoughtworks.gauge.test.git;

import static com.thoughtworks.gauge.test.common.GaugeProject.getCurrentProject;

import java.io.IOException;

import com.thoughtworks.gauge.Step;

public class Config {

    public enum GitConfig {
        HTTPS {
            @Override
            public String remoteOriginURL() {
                return "https://github.com/example-user/example-repo.git";
            }
        },
        SSH {
            @Override
            public String remoteOriginURL() {
                return "git@github.com:example-user/example-repo.git";
            }
        };

        public abstract String remoteOriginURL();
    }

    @Step("Add <type> Git config to project")
    public void addGitConfigToProject(String gitConfig) throws Exception {
        getCurrentProject().addGitConfig(GitConfig.valueOf(gitConfig));
    }

    @Step("Simulate Git detached HEAD")
    public void simulateGitDetachedHead() throws IOException {
        getCurrentProject().simulateGitDetachedHead();
    }

}
