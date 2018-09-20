Feature: Version

  As a software developer
  I need to know the sdkman version
  So that I can stay up to date

  Scenario: Show the current version of sdkman
    Given the internet is reachable
    And an initialised environment
    And the sdkman version is "6.0.0"
    When I enter "sdk version"
    Then I see "SDKMAN 6.0.0"
    And the exit code is 0
