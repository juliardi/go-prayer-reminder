# Go Prayer Reminder
Islamic Prayer Time Reminder written in Golang

# Requirements
To develop and build Go Prayer Reminder, you need to install the following tools :
- Go 1.8
- [Dep : Golang Dependency Management Tool](https://github.com/golang/dep)
- [UPX](https://upx.github.io/)
- [mpg123](https://sourceforge.net/projects/mpg123/)
- Azan Mp3 File

NOTE : Make sure that `mpg123` is accessible through command line

# Setup
### Installing Dependency
You can install the dependency using Dep by running the following command in the project directory :
````
$ dep ensure
````

### Create configuration file
Go Prayer Reminder stores the configuration in `config.env` file. You can create `config.env` file by copying and renaming `config.env.example` file and fill all the configuration key.
Here is an example of configuration for my hometown :
````
CITY_NAME="Ponorogo"
CITY_LAT="-8.00"
CITY_LONG="111.30"
CITY_TIMEZONE="7"
AZAN_FILENAME="path-to-mp3-file"
METHOD="MWL"
````
The METHOD setting, can be filled with one of these value :
- MWL
- ISNA
- Egypt
- Makkah
- Karachi
- Tehran
- Jafari

### Run / Build
You could run the program by running the following command in the project directory :
````
$ go run PrayerReminder.go
````

Or you could build it first by running the following command :
````
$ go build PrayerReminder.go
````

And then run the result executable :
````
$ ./PrayerReminder
````
