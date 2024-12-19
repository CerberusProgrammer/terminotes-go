@echo off
echo Building terminotes...
go build -o tn.exe ./src/main.go
if %errorlevel% neq 0 (
    echo Build failed.
    exit /b %errorlevel%
)
echo Build complete. The executable is tn.exe

REM Agregar la ruta del binario a PATH si no está ya presente
set "BIN_PATH=%~dp0"
echo Checking if %BIN_PATH% is already in PATH...
echo %PATH% | findstr /i /c:"%BIN_PATH%" >nul
if %errorlevel% neq 0 (
    echo Adding %BIN_PATH% to PATH...
    setx PATH "%PATH%;%BIN_PATH%"
    if %errorlevel% neq 0 (
        echo Failed to add %BIN_PATH% to PATH.
        exit /b %errorlevel%
    )
    echo %BIN_PATH% added to PATH successfully.
) else (
    echo %BIN_PATH% is already in PATH.
)