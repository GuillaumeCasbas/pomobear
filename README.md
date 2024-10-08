![Pomobear usage](/assets/pomobear-usage.png)

# Pomobear - The tool to be(ar) Effective

`pomobear` is a command line tool to help being focus & effective. It is based on the [Pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique). 

This project is a side-project and still in progress. Many things need to be added to fully apply the technique.

Of course, I could use one of the many excellent [similar app](https://github.com/topics/pomodoro) on GitHub, but building it is a great way of learning.

Advices, feedbacks and code reviews are welcomed.

## Documentation

For [installation options see below](#installation), for usage instructions [see the usage](#usage).

## Installation

### From release

Download the latest build from the [release page](https://github.com/GuillaumeCasbas/pomobear/releases/lastest)

### From source code

use `go build` and then move the executation somewhere (ex: `mv pomobear /usr/local/bin`).

## Usage

```
completion  Generate the autocompletion script for the specified shell
help        Help about any command
start       Start a pomodoro
status      Display the pomodoro remaining time (--raw to get the number of remaining seconds)
stop        Stop the running pomodoro
```

## Tmux integration


![Tmux integration example](/assets/pomobear-tmux.png)

The status countdown can be displayed inside the tmux status-bar with the following configuration :

```
set -g status-interval 1
set -g status-right "#(pomobear status)"

```
