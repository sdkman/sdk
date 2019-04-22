Feature: Selfupdate

  The selfupdate feature serves the purpose of upgrading sdkman itself. It does so by checking it's own version to
  the remote version that was previously pulled down by the `pull` command. If the versions differ, the `selfupdate`
  command will prompt the user for permission to upgrade.

  If the user chooses to upgrade, it will proceed by calling the selfupdate hook on the sdkman-hooks service. After
  this has completed, the local version file will also be updated. If the user chooses not to upgrade, the system will
  remain unchanged.

  In other words:
  $ sdk pull          #get the remotely available version
  $ sdk selfupcate    #compare the remote version to the local and

  local version file : `${SDKMAN_DIR}/var/version`
  remote version file: `${SDKMAN_DIR}/var/api.sdkman.io/version` (api.sdkman.io is the remote domain)

  Scenario: SDKMAN is out of date
    Given an initialised environment
    And the installed sdkman version is "5.7.2"
    And the pulled version state is "6.0.0"
    When I enter "sdk selfupdate"
    Then I see "SDKMAN 6.0.0 now available for installation..."
    And I see "Would you like to upgrade now? (Y/n):"
    Then I see "Successfully upgraded SDKMAN!"
    And the installed sdkman version was upgraded to "6.0.0"
    And the exit code is 0
