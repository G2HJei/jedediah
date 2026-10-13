# Jedediah

_Named for Jedediah Smith (1799–1831), fur trapper and pathfinder of the American West._

Jedediah tracks your job applications. It runs on your own computer and you use it in the browser:
where you applied, who you talked to, upcoming interviews, and what to do next and when. 

Your data stays on your machine as plain YAML files you can read and back up. 

There is no account and no cloud, and it works offline.

![Dashboard: overdue actions, follow-ups, the pipeline and the next two weeks](docs/screenshots/overview.png)

Every application in one table, which you can filter, sort, and search:

![Applications table](docs/screenshots/applications.png)

Everything about one application on one page: company, contacts, the posting, the hiring process,
interviews, questions and answers, compensation, and notes. The key dates and the next step are at
the top:

![Application page header: key dates, next step, priority and status](docs/screenshots/application-header.png)

![Job Application and Contacts cards](docs/screenshots/application-cards.png)

Click any field to edit it:

![The Interviews card, and the same card while editing](docs/screenshots/application-edit.png)

A board to drag applications from one status to the next:

![Board](docs/screenshots/board.png)

There is also a calendar (with an `.ics` export for Outlook or Google Calendar), a CV library, a
CSV export, and a ZIP backup.

## Run it

Jedediah can be run as a Docker container or as a Windows `.exe`.

### With Docker

Install [Docker Desktop](https://www.docker.com/products/docker-desktop/), then run this in a terminal:

```bash
docker run -d --name jedediah -p 127.0.0.1:8765:8765 -v "$HOME/Jedediah:/data" g2hjei/jedediah
```

Open http://localhost:8765. 

Optional arguments:

`$HOME/Jedediah` changes the data directory for the Jedediah data files.

`-e TZ=<timezone>` sets the timezone. This is important for the calendar to display the correct time. Defaults to `UTC`.

`--restart unless-stopped` Restarts Jedediah after a crash and whenever Docker starts..


### As a Windows executable

[Download the latest release](https://github.com/G2HJei/jedediah/releases/latest/download/jedediah.exe) or build the executable yourself following these steps:

1. Install [Go](https://go.dev/dl/) (1.27 or newer) and [Git](https://git-scm.com/downloads), then
run this in a terminal:

```bash
git clone https://github.com/G2HJei/jedediah.git
cd jedediah
go build -o jedediah.exe ./cmd/jedediah
```

2. Put `jedediah.exe` in any folder and run it. It opens Jedediah in your browser and keeps
your data in a `data` folder next to itself. To stop it, right-click the Jedediah icon in the tray
and choose **Quit**.

## More

[SPECS.md](SPECS.md) has the details: every feature, where your data lives and how to back it up,
the command-line flags, security, more Docker options, and how to develop Jedediah.

GPL-3.0, see [LICENSE](LICENSE).
