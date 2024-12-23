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

REM Create the installation directory if it doesn't exist
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

REM Copy the executable to the installation directory
copy /Y "%~dp0terminotes.exe" "%INSTALL_DIR%\terminotes.exe"
copy /Y "%INSTALL_DIR%\terminotes.exe" "%INSTALL_DIR%\tn.exe"

REM Check if the installation directory is already in the PATH
echo %PATH% | find /I "%INSTALL_DIR%" >nul
if %ERRORLEVEL% NEQ 0 (
    REM Add the installation directory to the system PATH
    setx PATH "%PATH%;%INSTALL_DIR%"
)

echo Terminotes has been installed successfully.
echo You can now run the application by typing "tn" in the terminal.

endlocal
pause