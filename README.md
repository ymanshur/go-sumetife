# Sumetife App
AccetByte Inc technical project-based test

### Build the app
Build the app using `make`, will return a (executable) file which name is sumetife.exe
```powershell
make build
```

### Running the app
Run `main.go` using `make`
```powershell
make run
```

If you already build the app, just run executable (.exe) file using following command
```powershell
.\sumetife.exe -d inputDirPath -t inputFileType --startTime inputStartTime --endTime inputEndTime
```
<small>Note: inputDirPath, inputFileType, inputStartTime, inputEndTime are required value</small>

Example command
```powershell
.\sumetife.exe -d data -t json --startTime 2022-01-01T07:00:20+07:00 --endTime 2022-01-02T00:00:00Z
```

#### Command line flag
| Flag | Description | Type |
| - | - | - |
| `-d` or `--directory` | The directory path, the directory contains single type of file, it can be csv or json | Required |
| `-t` or `--type` | The type of the input files, supported format: json and csv | Required |
| `--startTime` | The starting time to scan the data in the format of rfc3339, inclusive | Required |
| `--endTime` | The ending time to scan the data in the format of rfc3339, exclusive | Required |
| `--outputFileType` | The output type of the summary, supported value: json and yaml | Optional |
| `--outputFileName` | The output filename of summary | Optional |

### Testing the app
Complete (include function mode) run test
```powershell
make test
```

Add `mode=html` argument for html display
```powershell
make test mode=html
```