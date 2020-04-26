# go-simple-chart-project
simple golang chart project<br>
It's my university homework, which I decide to combine with code practice<br>
It consist of some boring math operation and some go code(that more intresting for me and for humanity(LOL))
# Program appearence 
![Programm appearence](https://i.ibb.co/gTnC4VT/image.png)
# Building for your platform from Windows OS
(this block in progress)<br>
This block was written primarily for me, because I needed to have this information somewhere
<h3> Windows </h3>
Just set up in root directory <a href="https://www.msys2.org/">https://www.msys2.org/</a> (necessary for fyne dependency)
<br><br>
After that clone repo to your and do (init params throw "set")
-with console<br>
CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -v
<br><br>

-without console(thanks to <a href="https://github.com/pymq">pymq</a>)<br>
CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui"
