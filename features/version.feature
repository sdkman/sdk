Feature: Version

  As a software developer
  I need to know the sdkman version
  So that I can stay up to date

  Scenario: Show the current version of sdkman
    Given the internet is reachable
    And the sdkman version is "3.2.1"
    When I enter "sdk version"
    Then I see "SDKMAN 3.2.1"
