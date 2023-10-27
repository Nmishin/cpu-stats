cpu-stats
=========
[![Lint](https://github.com/Nmishin/cpu-stats/actions/workflows/lint.yml/badge.svg)](https://github.com/Nmishin/cpu-stats/actions/workflows/lint.yml)

Application for check free memory and number of running processes.

After the application run, statistic available by address http://\<hostname\>:8080/metrics.

Where hostname - the hostname or IP address of the monitored server.

Limitations:
=========
Because of default tags generated by kvendingoldo/semver-action (with rc/ prefix) uncompatible with GoReleaser - need to add [RELEASE] prefix for each commit message. This will allow us to have a tag version without rc/ prefix, and build a new app version.
