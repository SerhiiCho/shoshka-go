![php revival](https://raw.githubusercontent.com/SerhiiCho/shoshka-go/master/storage/cover.jpg)

<p align="center">
    <a title="Latest push build on default branch: created">
        <img src="https://travis-ci.org/SerhiiCho/shoshka-go.svg?branch=master" alt="build:created">
    </a>
</p>

# Telegram bot

This program was created for running as a linux cron job, that makes checks every certain period of time for 2 things

* Checks for new posts on target site, and if there are new posts. It will notify a site administrator.
* Checks if there are any new errors in error_log on production server, if it is, notify a site administrator