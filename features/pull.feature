@wip
Feature: Pull

  As a Software Developer
  I want SDKMAN to be up-to-date
  So that I can access the latest SDKs

  https://github.com/sdkman/sdk/issues/4

  Versions of SDKMAN prior to 6.0.0 were very reliant on being connected to the internet. For most this has not
  been a problem, but for those with limited bandwidth or unreliable internet connectivity this has rendered the
  CLI unusable.

  One of the steps that was taken previously was to introduce an offline mode, allowing the CLI to continue working
  even when no internet connection was present. The intent was to cause the CLI to degrade gracefully when:
  a) no internet connectivity was available
  b) when the `sdk offline enable` command was issued

  Despite this being a noble effort, the implementation of this was troublesome due to the limitations of bash
  and curl. To date this functionality still has serious issues.

  This leads us to re-evaluating our options given a blank canvas. Looking at tools like Debian's Apt and DVCS systems
  like Git, they all have functions that explicitly pull down remote state. In the case of Apt we have an `apt update`,
  in the case of Git we have `git pull`.

  The proposal is to follow this line of thinking, preventing continual API access on every command. Instead, an
  explicit command such as `sdk pull` or `sdk update` should be issued in order to fetch remote state. If the sdk
  command has not been updated in while, it should warn the user that their local state is out of date.

