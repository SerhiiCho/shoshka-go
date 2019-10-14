<p align="center">
    <a href="https://travis-ci.org/SerhiiCho/shoshka-go">
        <img src="https://travis-ci.org/SerhiiCho/shoshka-go.svg?branch=master">
    </a>
    <a href="https://scrutinizer-ci.com/g/SerhiiCho/shoshka-go/">
        <img src="https://scrutinizer-ci.com/g/SerhiiCho/shoshka-go/badges/build.png?b=master">
    </a>
    <a href="https://scrutinizer-ci.com/g/SerhiiCho/shoshka-go/">
        <img src="https://scrutinizer-ci.com/g/SerhiiCho/shoshka-go/badges/quality-score.png?b=master">
    </a>
</p>

_This program is written for a specific case and is not helpful for others, but of course you can do whatever you want with it if you want_

## About

This program was created for running as a linux cron job, that makes checks every certain period of time for 3 things

* Checks for new posts on target site, and if there are new posts. It will notify a site administrator.
* Checks if there are any new errors in error_log on production server, if it is, notify a site administrator
* Pings given address from time to time to check if server is not down

## Program takes one of 3 arguments

```bash
./shoshka-go errors # Checks error log file in provided path, if there are new errors added to it it will send a telegram message to a certain chat.
```
```bash
./shoshka-go titles # Checks new photo reports at specific target site by parsing html and checking if there are new posts added. If new added, it takes the title, link and image url and sends it to a telegram chat.
```
```bash
./shoshka-go ping # Checks the state of the production server by using ping command. Sends 3 ping requests and if 1 of 3 requests is not successful, sends telegram message to a provided chat.
```