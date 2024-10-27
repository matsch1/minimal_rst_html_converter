# Minimal RST to HTML converter

I build a converter which does not need an API key or some web interface.
This is needed to work with text containing information you do not want to share with annyone.

## usage
### files
 `input.rst`: containing the restructered text to convert <br>
 `style.html`: containing html style decissions<br>
 `index.html`: output with converted text<br>
 `gotool`: amd64 executable build from go<br>
 `Taskfile.yml`: task automation to build and run the converter


Copy your rst text into input.rst and run `task run`and get the translated text in *index.html*.

### precondtions
tools nedded:
- go
- task