@echo off
setlocal

REM Check for administrative privileges
openfiles >nul 2>&1
if %errorlevel% neq 0 (
    echo Requesting administrative privileges...
    powershell -Command "Start-Process '%~f0' -Verb RunAs"
    exit /b
)

REM Define the installation directory
set "INSTALL_DIR=%ProgramFiles%\Terminotes"

REM Delete the executable from the installation directory
if exist "%INSTALL_DIR%\tn.exe" del "%INSTALL_DIR%\tn.exe"

REM Remove the installation directory if it is empty
if exist "%INSTALL_DIR%" rd "%INSTALL_DIR%"

REM Remove the installation directory from the system PATH
setlocal enabledelayedexpansion
set "NEW_PATH="
for %%i in ("%PATH:;=" "%") do (
    if /I "%%~i" neq "%INSTALL_DIR%" (
        if defined NEW_PATH (
            set "NEW_PATH=!NEW_PATH!;%%~i"
        ) else (
            set "NEW_PATH=%%~i"
        )
    )
)
endlocal & setx PATH "%NEW_PATH%"

echo Terminotes has been uninstalled successfully.
echo The command "tn" is no longer available.

endlocal
pause