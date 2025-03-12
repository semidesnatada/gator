# Gator: a CLI tool which aggregates RSS feeds.

## Built using Golang and postgresql.

### This project is functional but could do with some work. For example, it could be improved by:
* adding concurrent fetching and printing articles
* allowing selection of articles and launching a view of full article in browser
* better management of feeds (i.e. via feed name, rather than url)
* removal / improvement of the clunky user system
* quality of life features, such as help menus which explain the possible commands
* better / more consistent handling of multiple command line arguments.

### If you want to run this project on your machine, then clone this repo and build the project.
You will need to create a `.gatorconfig.json` config file in your home directory. This just needs to include the `db_url` string. This will depend on the name and port your database runs on. My database name can be found in the shell helper scripts included in this repo.