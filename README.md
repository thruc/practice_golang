# practice_golang


launch.json
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "unitfile",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}",
            "cwd": "${fileDirname}",
            "env": {},
            "args": []
        },
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
}
```