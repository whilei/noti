# Change Log

All notable changes to this project will be documented in this file. This
project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

## Added
* Tests.

## Fixed
* Setting `-t` or `-m` will now take precedence over utility name.

## Changed
* Slackbot icon is now a rocket.

## [2.1.0]

## Added

* `NOTI_SOUND_FAIL` tells Noti which sound to play for banner notifications
when a utility fails on OS X.
* `NOTI_SOUND` and `NOTI_SOUND_FAIL` can be set to `_mute` for silent
notifications.
* HipChat notifications.
* Pushover notifications.
* Noti Wiki.

## Changed

* On OS X, banner notifications will play different sounds depending on the
utility's success or failure, instead of the same sound for both.
* `NOTI_SLACK_DEST` no longer defaults to "#random". It must be manually set.

## [2.0.0]

## Added

* Speech notifications.
* Slack notifications.
* Optionally set default notification type(s) through `NOTI_DEFAULT` env var.
* Multi-notification support.
* Other configuration through environment variables.
* Contributing document.
* Change log.

## Changed

* On OS X, the notification sound must now be set in the environment variable,
`NOTI_SOUND`.
* On OS X, instead of AppleScript, banner notifications are triggered with
Object-C, which shows (nicer) Terminal icon.

## Removed

* `-f` flag for OS X. This caused unexpected behavior for people who use iTerm2.
* OS X-specific flags and usage text from Linux and FreeBSD help.

[Unreleased]: https://github.com/variadico/noti/compare/v2.1.0...dev
[2.1.0]: https://github.com/variadico/noti/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/variadico/noti/compare/v1.3.0...v2.0.0