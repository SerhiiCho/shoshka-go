<p align="center">
    <a title="Latest push build on default branch: created">
        <img src="https://travis-ci.org/SerhiiCho/shoshka-go.svg?branch=master" alt="build:created">
    </a>
</p>

```
This program is written for a specific case and is not helpful for others, but of course you can do whatever you want with it if you wan
```

## About

This program was created for running as a linux cron job, that makes checks every certain period of time for 3 things

* Checks for new posts on target site, and if there are new posts. It will notify a site administrator.
* Checks if there are any new errors in error_log on production server, if it is, notify a site administrator
* Pings given address from time to time to check if server is not down

## Program takes one of 3 arguments

```bash
./shoshka-go errors # for checking errors
```
```bash
./shoshka-go titles # for checking new photo reports
```
```bash
./shoshka-go ping # for checking the state of the production site by using ping
```