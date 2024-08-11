# Weather updates for developers!
A CLI application built with Go, created so you can check the weather no matter the directory you are working in.
## Demo
![cliweather](https://github.com/user-attachments/assets/4f1f609f-b7dd-4f7d-87cd-55b9f329888d)
## Installation Instructions:
#### Make sure to:
Install Go: https://go.dev/doc/install
### MacOS
1. Change current directory to weather
2. Build the executable:
```
go build -o weather
```
3. Move executable to bin:
```
sudo mv weather /usr/local/bin/
```
4. Double check PATH:
```
export PATH=$PATH:/usr/local/bin
```
5. Apply changes:
```
source ~/.zshrc
```
### How to use:
Simply call the app with 'weather', followed by a zipcode or city</br>
Examples:</br>
```weather columbus_ohio```</br>
```weather 10029```
### Windows
1. Change current directory to weather
2. Build the executable
```
go build -o weather.exe
```
3. Make custom directory for executable
```
mkdir C:\Users\{your-username-here}\bin
```
4. Move executable to directory
```
move weather.exe C:\Users\{your-username-here}\bin
```
5. Add directory to PATH manually
```
- Search for 'edit the system environment variables' or 'PATH'
- Click 'Environment Variables'
- Click 'Path' in User variables, then click Edit
- Create new environment variable, and insert the directory you made in step 4
```
You should be able to use the executable weather from any directory on your computer
### How to use:
Simply call the app with 'weather', followed by a zipcode or city</br>
Examples:</br>
```weather columbus_ohio```</br>
```weather 10029```
### Linux
-WIP
