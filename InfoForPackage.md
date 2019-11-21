# GoLangTutorial
In Go, it uses different method for using packaging from other languages like Java, Python, etc.
If you want to work on the Go packages, 
    1. Have src directory in the project's root directory
    2. export or set root directory as GOPATH value
    3. create a directory inside src directory and create a main file inside the directory
    4. write the Go source code
    5. come to the root directory
    6. Execute "go install <directory_name_inside_src>"
    7. Go will find the main package from the given directory that is inside the src
        and create a bin file from the src
    8. Execute the file to check whether the code is working or not
    9. To use the package from another file, create the executable binary file and import them using file       path