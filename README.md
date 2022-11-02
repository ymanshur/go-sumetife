# Sumetife App
AccetByte Inc technical project-based test

### Running the app

Running executable (.exe) file
```powershell
.\sumetife.exe -d inputDirPath -t inputFileType --startTime inputStartTime --endTime inputEndTime
```
Note: inputDirPath, inputFileType, inputStartTime, inputEndTime are required value

#### Command line flag
| Flag | Description | Type |
| - | - | - |
| `-d` or `--directory` | The directory path, the directory contains single type of file, it can be csv or json | Required |
| `-t` or `--type` | The type of the input files, supported format: json and csv | Required |
| `--startTime` | The starting time to scan the data in the format of rfc3339, inclusive | Required |
| `--endTime` | The ending time to scan the data in the format of rfc3339, exclusive | Required |
| `--outputFileType` | The output type of the summary, supported value: json and yaml | Optional |
| `--outputFileName` | The output filename of summary | Optional |