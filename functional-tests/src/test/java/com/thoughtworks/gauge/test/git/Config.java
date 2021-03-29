package com.thoughtworks.gauge.test.git;

import static com.thoughtworks.gauge.test.common.GaugeProject.getCurrentProject;

import java.io.File;
import com.thoughtworks.gauge.Step;

public class Config {

    public enum GitConfig {
        HTTPS, SSH;

        public File file() {
            return new File("src/test/resources/gitconfig/" + this.toString());
        }
    }

    @Step("Add <type> Git config to project")
    public void addGitConfigToProject(String gitConfig) throws Exception {
        getCurrentProject().addGitConfig(GitConfig.valueOf(gitConfig));
    }

}
